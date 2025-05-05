package main

import (
	"context"
	"strings"

	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	userpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	idempotencypostgres "github.com/Doremi203/couply/backend/auth/pkg/idempotency/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		userRepo := userpostgres.NewRepo(dbClient)

		passwordHasher := pswrd.NewDefaultHasher(
			salt.DefaultProvider{},
			argon.V2Provider{},
		)

		jwtTokenConfig := token.JWTConfig{}
		err = app.Config.ReadSection("jwt", &jwtTokenConfig)
		if err != nil {
			return err
		}

		tokenIssuer := token.NewJWTIssuer(jwtTokenConfig)

		registrationUseCase := usecase.NewRegistration(
			userRepo,
			passwordHasher,
			uuid.DefaultProvider{},
		)
		registrationService := grpc.NewRegistrationService(
			registrationUseCase,
			app.Log,
			postgres.NewProvider(dbClient.Pool),
			idempotencypostgres.NewRepo(dbClient),
		)

		loginUseCase := usecase.NewLogin(
			userRepo,
			passwordHasher,
			tokenIssuer,
		)
		loginService := grpc.NewLoginService(
			loginUseCase,
			app.Log,
		)

		app.RegisterGRPCServices(registrationService, loginService)
		app.AddGatewayHandlers(registrationService, loginService)

		return nil
	})
}
