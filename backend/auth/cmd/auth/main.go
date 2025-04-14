package main

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	"github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/idempotency/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"strings"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		app.AddGatewayOptions(
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "idempotency-key":
					return s, true
				default:
					return runtime.DefaultHeaderMatcher(s)
				}
			}),
		)

		dbConfig := postgres.Config{}
		err := app.Config.ReadSection("database", &dbConfig)
		if err != nil {
			return err
		}

		// temporal
		app.Log.Info("TEST LOG", "db_user", dbConfig.User)

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		userRepo := userpostgres.NewRepo(dbClient)

		registrationUseCase := usecase.NewRegistration(
			userRepo,
			pswrd.NewDefaultHasher(
				salt.DefaultProvider{},
				argon.V2Provider{},
			),
			uuid.DefaultProvider{},
		)

		registrationService := grpc.NewRegistrationService(
			registrationUseCase,
			app.Log,
			postgres.NewProvider(dbClient.Pool),
			idempotencypostgres.NewRepo(dbClient),
		)

		app.RegisterGRPCServices(registrationService)
		app.AddGatewayHandlers(registrationService)

		return nil
	})
}
