package main

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	userrepo "github.com/Doremi203/couply/backend/auth/internal/repo/user"
	"github.com/Doremi203/couply/backend/auth/internal/services/hash"
	"github.com/Doremi203/couply/backend/auth/internal/usecase/registration"
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

		userRepo := userrepo.NewPostgresRepo(pgPool)

		registerWithCredentialsUsecase := registration.Basic{
			UserRepository: userRepo,
			Hasher:         hash.NewArgon(app.Log),
			UIDGenerator:   &user.UUIDV7BasedUIDGenerator{},
		}

		app.RegisterGRPCService(grpc.NewRegistrationService(
			registerWithCredentialsUsecase,
			app.Log,
		))

		return nil
	})
}
