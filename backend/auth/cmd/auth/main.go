package main

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/gen/api/registration"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	"github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/services/hash"
	registrationUC "github.com/Doremi203/couply/backend/auth/internal/usecase/registration"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		dbConfig := postgres.Config{}

		err := app.Config.ReadSection("database", &dbConfig)
		if err != nil {
			return err
		}

		pgPool, err := pgxpool.New(ctx, dbConfig.ConnectionString())
		if err != nil {
			return errors.Wrap(err, "could create connection pool for db")
		}

		userRepo := userpostgres.NewRepo(pgPool)

		registerWithCredentialsUsecase := registrationUC.NewUseCase(
			userRepo,
			hash.NewArgon(app.Log),
			&user.UUIDV7BasedUIDGenerator{},
		)

		registrationService := grpc.NewRegistrationService(
			registerWithCredentialsUsecase,
			app.Log,
		)

		app.RegisterGRPCService(registrationService)

		err = app.RegisterGatewayHandler(registration.RegisterRegistrationHandlerFromEndpoint)
		if err != nil {
			return errors.Wrap(err, "failed to register gateway handler for registration")
		}

		return nil
	})
}
