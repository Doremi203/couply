package main

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/client/matcher"
	postgrestx "github.com/Doremi203/couply/backend/payments/internal/storage"
	payment_facade "github.com/Doremi203/couply/backend/payments/internal/storage/payment/facade"
	postgres3 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	subscription_facade "github.com/Doremi203/couply/backend/payments/internal/storage/subscription/facade"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	payment_service "github.com/Doremi203/couply/backend/payments/internal/app/payment-service"
	subscription_service "github.com/Doremi203/couply/backend/payments/internal/app/subscription-service"
	payment_usecase "github.com/Doremi203/couply/backend/payments/internal/usecase/payment-service"
	"github.com/Doremi203/couply/backend/payments/internal/usecase/payment-service/mock_gateway"
	subscription_usecase "github.com/Doremi203/couply/backend/payments/internal/usecase/subscription-service"
	"github.com/Doremi203/couply/backend/payments/internal/usecase/updater"
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

		var matcherConfig struct {
			Address string `yaml:"address"`
		}
		if err = app.Config.ReadSection("matcher", &matcherConfig); err != nil {
			return errors.WrapFail(err, "read user service config")
		}

		userServiceClient, conn, err := matcher.NewClient(matcherConfig.Address)
		if err != nil {
			return errors.WrapFail(err, "create user service client")
		}
		app.AddCloser(conn.Close)

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		txManager := postgrestx.NewTxManager(dbClient)

		subRepo := postgres2.NewPgStorageSubscription(txManager)
		payRepo := postgres3.NewPgStoragePayment(txManager)

		subFacade := subscription_facade.NewStorageFacadeSubscription(txManager, subRepo, payRepo)
		payFacade := payment_facade.NewStorageFacadePayment(txManager, payRepo, subRepo)

		gateway := mock_gateway.NewMockGateway()

		asyncUpdater := updater.NewUpdater(payFacade, subFacade, gateway, userServiceClient, app.Log)

		updaterCtx, updaterCancel := context.WithCancel(ctx)
		app.AddCloser(func() error {
			updaterCancel()
			return nil
		})

		go asyncUpdater.StartPaymentStatusUpdater(updaterCtx, 30*time.Second)
		go asyncUpdater.StartSubscriptionStatusUpdater(updaterCtx, 1*time.Hour)

		subUseCase := subscription_usecase.NewUseCase(subFacade)
		payUseCase := payment_usecase.NewUseCase(payFacade, gateway, asyncUpdater)

		subService := subscription_service.NewImplementation(subUseCase)
		payService := payment_service.NewImplementation(payUseCase)

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
