package main

import (
	"context"
	"strings"

	phoneconfirmgrpc "github.com/Doremi203/couply/backend/auth/gen/api/phone-confirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	phoneconfirmpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/phoneconfirm/postgres"
	userpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	idempotencypostgres "github.com/Doremi203/couply/backend/auth/pkg/idempotency/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/sms"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		app.AddGatewayOptions(
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "idempotency-key", "user-token":
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

		jwtTokenConfig := token.JWTConfig{}
		err = app.Config.ReadSection("jwt", &jwtTokenConfig)
		if err != nil {
			return err
		}

		pkgTokenConfig := tokenpkg.Config{}
		err = app.Config.ReadSection("token", &pkgTokenConfig)
		if err != nil {
			return err
		}

		phoneConfirmationConfig := phoneconfirm.Config{}
		err = app.Config.ReadSection("phone-confirmation", &phoneConfirmationConfig)
		if err != nil {
			return err
		}

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		userRepo := userpostgres.NewRepo(dbClient)

		hashProvider := hash.NewDefaultProvider(
			salt.DefaultProvider{},
			argon.V2Provider{},
		)

		tokenIssuer, err := token.NewJWTIssuer(jwtTokenConfig)
		if err != nil {
			return err
		}

		var uuidProvider uuid.DefaultProvider

		registrationUseCase := usecase.NewRegistration(
			userRepo,
			hashProvider,
			uuidProvider,
		)
		registrationService := grpc.NewRegistrationService(
			registrationUseCase,
			app.Log,
			postgres.NewProvider(dbClient.Pool),
			idempotencypostgres.NewRepo(dbClient),
		)

		loginUseCase := usecase.NewLogin(
			userRepo,
			hashProvider,
			tokenIssuer,
		)
		loginService := grpc.NewLoginService(
			loginUseCase,
			app.Log,
		)

		phoneConfirmationUseCase := usecase.NewPhoneConfirmation(
			sms.TestSender{
				Logger: app.Log,
			},
			phoneconfirm.NewDigitCodeGenerator(phoneConfirmationConfig),
			hashProvider,
			phoneconfirmpostgres.NewRepo(dbClient),
		)
		phoneConfirmationService := grpc.NewPhoneConfirmationService(
			phoneConfirmationUseCase,
			app.Log,
		)

		tokenProvider := tokenpkg.NewJWTProvider(pkgTokenConfig)

		app.AddGRPCUnaryInterceptor(
			tokenpkg.NewUnaryTokenInterceptor(
				tokenProvider,
				app.Log,
				phoneconfirmgrpc.PhoneConfirmation_SendCodeV1_FullMethodName,
			),
		)
		app.RegisterGRPCServices(
			registrationService,
			loginService,
			phoneConfirmationService,
		)
		app.AddGatewayHandlers(
			registrationService,
			loginService,
			phoneConfirmationService,
		)

		return nil
	})
}
