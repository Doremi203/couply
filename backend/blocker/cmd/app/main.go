package main

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	blocker_service "github.com/Doremi203/couply/backend/blocker/internal/app/blocker-service"
	telegram_client "github.com/Doremi203/couply/backend/blocker/internal/client/telegram"
	"github.com/Doremi203/couply/backend/blocker/internal/client/user"
	blocker_usecase "github.com/Doremi203/couply/backend/blocker/internal/usecase/blocker-service"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		var telegramConfig struct {
			Token       string `secret:"telegram-token"`
			AdminChatID int64  `secret:"telegram-admin-chat-id"`
		}
		err := app.Config.ReadSection("telegram", &telegramConfig)
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

		blockUseCase := blocker_usecase.NewUseCase(
			userServiceClient,
			bot,
			app.Log,
		)

		blockService := blocker_service.NewImplementation(
			blockUseCase,
			app.Log,
		)

		app.RegisterGRPCServices(
			blockService,
		)

		app.AddGatewayHandlers(
			blockService,
		)

		app.AddBackgroundProcess(func(ctx context.Context) error {
			bot.StartCallbackHandler(func(userID string, userToken string) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				userFromClient, err := userServiceClient.GetUserByIDV1(ctx, userID, userToken)
				if err != nil {
					app.Log.Infof("Failed to get user: %v", err)
					return
				}

				userFromClient.IsBlocked = true

				if err := userServiceClient.UpdateUserByIDV1(
					ctx,
					userFromClient,
					userToken,
				); err != nil {
					app.Log.Infof("Failed to block user: %v", err)
				} else {
					app.Log.Infof("User %s blocked successfully", userID)
				}
			})
			return nil
		})

		return nil
	})
}
