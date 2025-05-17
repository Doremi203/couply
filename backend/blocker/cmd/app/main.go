package main

import (
	"context"

	postgrespkg "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	blockerservicegrpc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/facade"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/postgres/blocker"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	blocker_service "github.com/Doremi203/couply/backend/blocker/internal/app/blocker-service"
	telegram_client "github.com/Doremi203/couply/backend/blocker/internal/client/telegram"
	"github.com/Doremi203/couply/backend/blocker/internal/client/user"
	blocker_usecase "github.com/Doremi203/couply/backend/blocker/internal/usecase/blocker-service"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		dbConfig := postgrespkg.Config{}
		err := app.Config.ReadSection("database", &dbConfig)
		if err != nil {
			return err
		}

		dbClient, err := postgrespkg.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		var telegramConfig struct {
			Token       string `secret:"telegram-token"`
			AdminChatID int64  `secret:"telegram-admin-chat-id"`
		}
		err = app.Config.ReadSection("telegram", &telegramConfig)
		if err != nil {
			return errors.WrapFail(err, "read telegram config")
		}

		tokenConfig := token.Config{}
		err = app.Config.ReadSection("user-token", &tokenConfig)
		if err != nil {
			return err
		}

		var userServiceConfig struct {
			Address string `yaml:"address"`
		}
		if err := app.Config.ReadSection("user_service", &userServiceConfig); err != nil {
			return errors.WrapFail(err, "read user service config")
		}

		userServiceClient, conn, err := user.NewClient(userServiceConfig.Address)
		if err != nil {
			return errors.WrapFail(err, "create user service client")
		}
		app.AddCloser(conn.Close)

		bot, err := telegram_client.NewBotClient(telegramConfig.Token, telegramConfig.AdminChatID)
		if err != nil {
			return errors.WrapFail(err, "create telegram bot")
		}

		txManager := postgres.NewTxManager(dbClient)
		blockerStorage := blocker.NewPgStorageBlocker(txManager)
		blockerFacade := facade.NewStorageFacadeBlocker(txManager, blockerStorage)

		blockUseCase := blocker_usecase.NewUseCase(
			userServiceClient,
			bot,
			blockerFacade,
			app.Log,
		)

		blockService := blocker_service.NewImplementation(
			blockUseCase,
			app.Log,
		)

		app.AddGRPCUnaryInterceptor(
			token.NewUnaryTokenInterceptor(
				token.NewJWTProvider(tokenConfig),
				app.Log,
				blockerservicegrpc.BlockerService_GetBlockInfoV1_FullMethodName,
			),
		)
		app.RegisterGRPCServices(
			blockService,
		)
		app.AddGatewayHandlers(
			blockService,
		)

		initBotHandlers(app, userServiceClient, blockerFacade, bot)

		return nil
	})
}

func initBotHandlers(
	app *webapp.App,
	userClient *user.Client,
	blockerFacade *facade.StorageFacadeBlocker,
	bot *telegram_client.BotClient,
) {
	botHandler := telegram_client.NewBotHandler(userClient, blockerFacade, bot, app.Log)

	app.AddBackgroundProcess(func(ctx context.Context) error {
		botHandler.SetupRoutes()
		bot.StartCallbackHandler()
		return nil
	})
}
