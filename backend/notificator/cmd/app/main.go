package main

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	pushgrpc "github.com/Doremi203/couply/backend/notificator/gen/api/push"
	"github.com/Doremi203/couply/backend/notificator/internal/grpc"
	pushpostgres "github.com/Doremi203/couply/backend/notificator/internal/repo/push/postgres"
	"github.com/Doremi203/couply/backend/notificator/internal/usecase"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		dbConfig := postgres.Config{}
		err := app.Config.ReadSection("database", &dbConfig)
		if err != nil {
			return err
		}

		pkgTokenConfig := tokenpkg.Config{}
		err = app.Config.ReadSection("user-token", &pkgTokenConfig)
		if err != nil {
			return err
		}

		webPushConfig := struct {
			VapidPublicKey  string
			VapidPrivateKey string
		}{}
		err = app.Config.ReadSection("web-push", &webPushConfig)
		if err != nil {
			return err
		}

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		pushRepo := pushpostgres.NewRepo(dbClient)

		pushSubscriptionUseCase := usecase.NewPushSubscription(
			pushRepo,
		)
		pushSubscriptionService := grpc.NewPushSubscriptionService(
			pushSubscriptionUseCase,
			app.Log,
		)

		tokenProvider := tokenpkg.NewJWTProvider(pkgTokenConfig)

		app.AddGRPCUnaryInterceptor(
			tokenpkg.NewUnaryTokenInterceptor(
				tokenProvider,
				app.Log,
				pushgrpc.Subscription_SubscribeV1_FullMethodName,
				pushgrpc.Subscription_UnsubscribeV1_FullMethodName,
			),
		)
		app.RegisterGRPCServices(
			pushSubscriptionService,
		)
		app.AddGatewayHandlers(
			pushSubscriptionService,
		)

		return nil
	})
}
