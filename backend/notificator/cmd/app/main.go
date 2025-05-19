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
	"github.com/SherClockHolmes/webpush-go"
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
			VapidPublicKey  string `secret:"web-push-vapid-public-key"`
			VapidPrivateKey string `secret:"web-push-vapid-private-key"`
			Subscriber      string `secret:"web-push-subscriber"`
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

		pushSenderUseCase := usecase.NewPushSender(
			&webpush.Options{
				Subscriber:      webPushConfig.Subscriber,
				TTL:             60,
				VAPIDPublicKey:  webPushConfig.VapidPublicKey,
				VAPIDPrivateKey: webPushConfig.VapidPrivateKey,
			},
			pushRepo,
			app.Log,
		)

		pushAdminService := grpc.NewAdminService(
			pushRepo,
			pushSenderUseCase,
			app.Log,
		)

		tokenProvider := tokenpkg.NewJWTProvider(pkgTokenConfig)

		app.AddAPIKeyProtectedEndpoints(pushgrpc.Admin_SendPushV1_FullMethodName)
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
			pushAdminService,
		)
		app.AddGatewayHandlers(
			pushSubscriptionService,
			pushAdminService,
		)

		return nil
	})
}
