package main

import (
	"context"

	"github.com/Doremi203/couply/backend/common/libs/sqs"
	"github.com/Doremi203/couply/backend/matcher/gen/api/messages"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	matching_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/matching/facade"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"
	search_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/search/facade"
	postgres4 "github.com/Doremi203/couply/backend/matcher/internal/storage/search/postgres"
	user_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/user/facade"
	postgres3 "github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	postgrespkg "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	matchingservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	searchservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	matching_service "github.com/Doremi203/couply/backend/matcher/internal/app/matching-service"
	search_service "github.com/Doremi203/couply/backend/matcher/internal/app/search-service"
	user_service "github.com/Doremi203/couply/backend/matcher/internal/app/user-service"
	user_domain "github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	matching_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/matching-service"
	search_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/search-service"
	user_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/user-service"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
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

		sqsConfig := sqs.Config{}
		err = app.Config.ReadSection("sqs", &sqsConfig)
		if err != nil {
			return err
		}

		sqsClient, err := sqs.New[*messages.MatcherEvent](sqsConfig)
		if err != nil {
			return errors.WrapFail(err, "create sqs client")
		}

		txManager := storage.NewTxManager(dbClient)
		pgStorageUser := postgres3.NewPgStorageUser(txManager)
		storageFacadeUser := user_service_facade.NewStorageFacadeUser(txManager, pgStorageUser)
		useCaseUserService := user_service_usecase.NewUseCase(photoURLGenerator, storageFacadeUser)
		implUserService := user_service.NewImplementation(useCaseUserService)

		pgStorageMatching := postgres2.NewPgStorageMatching(txManager)
		storageFacadeMatching := matching_service_facade.NewStorageFacadeMatching(txManager, pgStorageMatching)
		useCaseMatchingService := matching_service_usecase.NewUseCase(storageFacadeMatching, sqsClient)
		implMatchingService := matching_service.NewImplementation(useCaseMatchingService)

		pgStorageSearch := postgres4.NewPgStorageSearch(txManager)
		storageFacadeSearch := search_service_facade.NewStorageFacadeSearch(txManager, pgStorageSearch, pgStorageUser)
		useCaseSearchService := search_service_usecase.NewUseCase(storageFacadeSearch, photoURLGenerator, app.Log)
		implSearchService := search_service.NewImplementation(useCaseSearchService, photoURLGenerator)

		app.AddGRPCUnaryInterceptor(
			token.NewUnaryTokenInterceptor(
				token.NewJWTProvider(tokenConfig),
				app.Log,
				userservicegrpc.UserService_CreateUserV1_FullMethodName,
				userservicegrpc.UserService_UpdateUserV1_FullMethodName,
				userservicegrpc.UserService_DeleteUserV1_FullMethodName,
				userservicegrpc.UserService_GetUserV1_FullMethodName,
				userservicegrpc.UserService_GetUsersV1_FullMethodName,
				userservicegrpc.UserService_ConfirmPhotosUploadV1_FullMethodName,
				searchservicegrpc.SearchService_CreateFilterV1_FullMethodName,
				searchservicegrpc.SearchService_UpdateFilterV1_FullMethodName,
				searchservicegrpc.SearchService_GetFilterV1_FullMethodName,
				searchservicegrpc.SearchService_SearchUsersV1_FullMethodName,
				searchservicegrpc.SearchService_AddViewV1_FullMethodName,
				matchingservicegrpc.MatchingService_LikeUserV1_FullMethodName,
				matchingservicegrpc.MatchingService_DislikeUserV1_FullMethodName,
				matchingservicegrpc.MatchingService_DeleteMatchV1_FullMethodName,
				matchingservicegrpc.MatchingService_FetchMatchesUserIDsV1_FullMethodName,
				matchingservicegrpc.MatchingService_FetchIncomingLikesV1_FullMethodName,
				matchingservicegrpc.MatchingService_FetchOutgoingLikesV1_FullMethodName,
			),
		)
		app.RegisterGRPCServices(implUserService, implMatchingService, implSearchService)
		app.AddGatewayHandlers(implUserService, implMatchingService, implSearchService)

		return nil
	})
}
