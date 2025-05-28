package main

import (
	"context"
	"time"

	telegram5 "github.com/Doremi203/couply/backend/auth/gen/api/telegram"
	telegram4 "github.com/Doremi203/couply/backend/auth/internal/grpc/telegram"
	telegram2 "github.com/Doremi203/couply/backend/auth/internal/repo/telegram"
	telegram3 "github.com/Doremi203/couply/backend/auth/internal/usecase/telegram"

	phoneconfirmgrpc "github.com/Doremi203/couply/backend/auth/gen/api/phone-confirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm/senders/fallback"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm/senders/smsru"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm/senders/telegram"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/grpc"
	"github.com/Doremi203/couply/backend/auth/internal/grpc/login"
	tokengrpc "github.com/Doremi203/couply/backend/auth/internal/grpc/token"
	phoneconfirmpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/phoneconfirm/postgres"
	tokenpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/token/postgres"
	userpostgres "github.com/Doremi203/couply/backend/auth/internal/repo/user/postgres"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	loginUC "github.com/Doremi203/couply/backend/auth/internal/usecase/login"
	tokenUC "github.com/Doremi203/couply/backend/auth/internal/usecase/token"
	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	idempotencypostgres "github.com/Doremi203/couply/backend/auth/pkg/idempotency/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
	tokenpkg "github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	"github.com/Doremi203/couply/backend/common/valkey"
)

func main() { //nolint:gocognit
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		valkeyRateLimiterConfig := valkey.RateLimiterConfig{}
		err := app.Config.ReadSection("valkey-rate-limiter", &valkeyRateLimiterConfig)
		if err != nil {
			return err
		}

		dbConfig := postgres.Config{}
		err = app.Config.ReadSection("database", &dbConfig)
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

		telegramSenderConfig := telegram.Config{}
		err = app.Config.ReadSection("telegram-sender", &telegramSenderConfig)
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

		vkOAuthConfig := oauth.VKConfig{}
		err = app.Config.ReadSection("oauth-vk", &vkOAuthConfig)
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
		tokenRepo := tokenpostgres.NewRepo(dbClient)
		telegramRepo := telegram2.NewRepo(dbClient)

		hashProvider := hash.NewDefaultProvider(
			salt.DefaultProvider{},
			argon.V2Provider{},
		)

		providerFactory := oauth.NewProviderFactory(yandexOAuthConfig, vkOAuthConfig)
		timeProvider := timeprovider.ProviderFunc(time.Now)

		tokenIssuer, err := token.NewJWTIssuer(jwtTokenConfig, tokenRepo, timeProvider)
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
			fallback.NewSender(
				telegram.NewSender(telegramSenderConfig),
				smsru.NewSender(
					smsruSenderConfig,
					app.HTTPClient(),
				),
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

		telegramUseCase := telegram3.NewUseCase(telegramRepo, txProvider)
		telegramService := telegram4.NewGRPCService(app.Log, telegramUseCase)

		tokenProvider := tokenpkg.NewJWTProvider(pkgTokenConfig)

		tokenUseCase := tokenUC.NewUseCase(tokenRepo, tokenIssuer, timeProvider)
		tokenService := tokengrpc.NewGRPCService(app.Log, tokenUseCase)

		if !valkeyRateLimiterConfig.Disabled {
			valkeyRateLimiter, err := valkey.NewValkeyRateLimiter(valkeyRateLimiterConfig)
			if err != nil {
				return errors.WrapFail(err, "create valkey rate limiter")
			}
			app.SetRateLimiter(valkeyRateLimiter)
		}

		app.AddGRPCUnaryInterceptor(
			tokenpkg.NewUnaryTokenInterceptor(
				tokenProvider,
				app.Log,
				phoneconfirmgrpc.PhoneConfirmation_SendCodeV1_FullMethodName,
				phoneconfirmgrpc.PhoneConfirmation_ConfirmV1_FullMethodName,
				telegram5.TelegramData_SetTelegramDataV1_FullMethodName,
				telegram5.TelegramData_GetTelegramDataV1_FullMethodName,
			),
		)
		app.RegisterGRPCServices(
			registrationService,
			loginService,
			phoneConfirmationService,
			tokenService,
			telegramService,
		)
		app.AddGatewayHandlers(
			registrationService,
			loginService,
			phoneConfirmationService,
			tokenService,
			telegramService,
		)

		return nil
	})
}
