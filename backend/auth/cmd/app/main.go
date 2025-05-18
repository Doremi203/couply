package main

import (
	"context"

	phoneconfirmgrpc "github.com/Doremi203/couply/backend/auth/gen/api/phone-confirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	"github.com/Doremi203/couply/backend/auth/internal/grpc/login"
	phoneconfirmpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/phoneconfirm/postgres"
	userpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	loginUC "github.com/Doremi203/couply/backend/auth/internal/usecase/login"
	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	idempotencypostgres "github.com/Doremi203/couply/backend/auth/pkg/idempotency/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/sms/smsru"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
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

		smsruSenderConfig := smsru.Config{}
		err = app.Config.ReadSection("smsru", &smsruSenderConfig)
		if err != nil {
			return err
		}

		oauthLoginConfig := login.Config{}
		err = app.Config.ReadSection("oauth", &oauthLoginConfig)
		if err != nil {
			return err
		}

		yandexOAuthConfig := oauth.YandexConfig{}
		err = app.Config.ReadSection("oauth-yandex", &yandexOAuthConfig)
		if err != nil {
			return err
		}

		dbClient, err := postgres.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		userRepo := userpostgres.NewRepo(dbClient)
		oauthAccountRepo := userpostgres.NewOAuthAccountRepo(dbClient)

		hashProvider := hash.NewDefaultProvider(
			salt.DefaultProvider{},
			argon.V2Provider{},
		)

		providerFactory := oauth.NewProviderFactory(yandexOAuthConfig)

		tokenIssuer, err := token.NewJWTIssuer(jwtTokenConfig)
		if err != nil {
			return err
		}

		var uuidProvider uuid.DefaultProvider
		txProvider := postgres.NewTxProvider(dbClient.Pool)

		registrationUseCase := usecase.NewRegistration(
			userRepo,
			hashProvider,
			uuidProvider,
		)
		registrationService := grpc.NewRegistrationService(
			registrationUseCase,
			app.Log,
			txProvider,
			idempotencypostgres.NewRepo(dbClient),
		)

		loginUseCase := loginUC.NewUseCase(
			userRepo,
			oauthAccountRepo,
			providerFactory,
			hashProvider,
			tokenIssuer,
			txProvider,
			app.Log,
			uuidProvider,
		)
		loginService := login.NewGRPCService(
			loginUseCase,
			oauthLoginConfig,
			app.Log,
		)

		phoneConfirmationUseCase := usecase.NewPhoneConfirmation(
			smsru.NewSender(
				smsruSenderConfig,
				app.HTTPClient(),
				app.Log,
			),
			phoneconfirm.NewDigitCodeGenerator(phoneConfirmationConfig),
			hashProvider,
			phoneconfirmpostgres.NewRepo(dbClient),
			userRepo,
			txProvider,
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
				phoneconfirmgrpc.PhoneConfirmation_ConfirmV1_FullMethodName,
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
