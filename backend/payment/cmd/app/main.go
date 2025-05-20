package main

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	payment_service3 "github.com/Doremi203/couply/backend/payment/internal/app/payment-service"
	subscription_service3 "github.com/Doremi203/couply/backend/payment/internal/app/subscription-service"
	payment_service "github.com/Doremi203/couply/backend/payment/internal/storage/facade/payment-service"
	subscription_service "github.com/Doremi203/couply/backend/payment/internal/storage/facade/subscription-service"
	postgres2 "github.com/Doremi203/couply/backend/payment/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres/payment"
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres/subscription"
	payment_service2 "github.com/Doremi203/couply/backend/payment/internal/usecase/payment-service"
	"github.com/Doremi203/couply/backend/payment/internal/usecase/payment-service/mock_gateway"
	subscription_service2 "github.com/Doremi203/couply/backend/payment/internal/usecase/subscription-service"
	updater2 "github.com/Doremi203/couply/backend/payment/internal/usecase/updater"
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

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		txManager := postgres2.NewTxManager(dbClient)

		subRepo := subscription.NewPgStorageSubscription(txManager)
		payRepo := payment.NewPgStoragePayment(txManager)

		subFacade := subscription_service.NewStorageFacadeSubscription(txManager, subRepo, payRepo)
		payFacade := payment_service.NewStorageFacadePayment(txManager, payRepo)

		gateway := mock_gateway.NewMockGateway()

		updater := updater2.NewUpdater(payFacade, subFacade, gateway, app.Log)

		updaterCtx, updaterCancel := context.WithCancel(ctx)
		app.AddCloser(func() error {
			updaterCancel()
			return nil
		})

		go updater.StartPaymentStatusUpdater(updaterCtx, 30*time.Second)
		go updater.StartSubscriptionStatusUpdater(updaterCtx, 1*time.Hour)

		subUseCase := subscription_service2.NewUseCase(subFacade)
		payUseCase := payment_service2.NewUseCase(payFacade, gateway, updater)

		subService := subscription_service3.NewImplementation(app.Log, subUseCase)
		payService := payment_service3.NewImplementation(app.Log, payUseCase)

		tokenProvider := tokenpkg.NewJWTProvider(pkgTokenConfig)

		app.AddGRPCUnaryInterceptor(
			tokenpkg.NewUnaryTokenInterceptor(
				tokenProvider,
				app.Log,
				tokenpkg.InterceptAllMethodsOption,
			),
		)
		app.RegisterGRPCServices(
			subService,
			payService,
		)
		app.AddGatewayHandlers(
			subService,
			payService,
		)

		return nil
	})
}
