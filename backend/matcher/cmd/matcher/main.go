package main

import (
	"context"
	"strings"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	postgrespkg "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	matching_service "github.com/Doremi203/couply/backend/matcher/internal/app/matching-service"
	search_service "github.com/Doremi203/couply/backend/matcher/internal/app/search-service"
	user_service "github.com/Doremi203/couply/backend/matcher/internal/app/user-service"
	user_domain "github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	matching_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/matching-service"
	search_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/search-service"
	user_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/user-service"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/search"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/user"
	matching_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/matching-service"
	search_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/search-service"
	user_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/user-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		app.AddGatewayOptions(
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "x-api-key", "user-token":
					return s, true
				default:
					return runtime.DefaultHeaderMatcher(s)
				}
			}),
		)

		dbConfig := postgrespkg.Config{}
		err := app.Config.ReadSection("database", &dbConfig)
		if err != nil {
			return err
		}

		tokenConfig := token.Config{}
		err = app.Config.ReadSection("user-token", &tokenConfig)
		if err != nil {
			return err
		}

		dbClient, err := postgrespkg.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		s3Config := struct {
			Endpoint  string
			AccessKey string `secret:"aws-access-key"`
			SecretKey string `secret:"aws-secret-key"`
			Bucket    string
			Secure    bool
		}{}

		err = app.Config.ReadSection("s3", &s3Config)
		if err != nil {
			return err
		}

		s3Client, err := minio.New(s3Config.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(s3Config.AccessKey, s3Config.SecretKey, ""),
			Secure: s3Config.Secure,
		})
		if err != nil {
			return errors.WrapFail(err, "create s3 client")
		}

		photoURLGenerator := user_domain.NewObjectStoragePhotoURLGenerator(s3Client, s3Config.Bucket)

		txManager := postgres.NewTxManager(dbClient)
		pgStorageUser := user.NewPgStorageUser(txManager)
		storageFacadeUser := user_service_facade.NewStorageFacadeUser(txManager, pgStorageUser)
		useCaseUserService := user_service_usecase.NewUseCase(photoURLGenerator, storageFacadeUser)
		implUserService := user_service.NewImplementation(app.Log, useCaseUserService)

		pgStorageMatching := matching.NewPgStorageMatching(txManager)
		storageFacadeMatching := matching_service_facade.NewStorageFacadeMatching(txManager, pgStorageMatching)
		useCaseMatchingService := matching_service_usecase.NewUseCase(storageFacadeMatching)
		implMatchingService := matching_service.NewImplementation(app.Log, useCaseMatchingService)

		pgStorageSearch := search.NewPgStorageSearch(txManager)
		storageFacadeSearch := search_service_facade.NewStorageFacadeSearch(txManager, pgStorageSearch, pgStorageUser)
		useCaseSearchService := search_service_usecase.NewUseCase(storageFacadeSearch)
		implSearchService := search_service.NewImplementation(app.Log, useCaseSearchService)

		app.AddGRPCUnaryInterceptor(
			token.NewUnaryTokenInterceptor(
				token.NewJWTProvider(tokenConfig),
				app.Log,
				token.InterceptAllMethodsOption,
			),
		)
		app.RegisterGRPCServices(implUserService, implMatchingService, implSearchService)
		app.AddGatewayHandlers(implUserService, implMatchingService, implSearchService)

		return nil
	})
}
