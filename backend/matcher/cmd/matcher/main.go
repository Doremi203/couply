package main

import (
	"context"
	"strings"

	search_service "github.com/Doremi203/couply/backend/matcher/internal/app/search-service"
	search_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/search-service"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/search"
	search_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/search-service"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	postgrespkg "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	matching_service "github.com/Doremi203/couply/backend/matcher/internal/app/matching-service"
	user_service "github.com/Doremi203/couply/backend/matcher/internal/app/user-service"
	matching_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/matching-service"
	user_service_facade "github.com/Doremi203/couply/backend/matcher/internal/storage/facade/user-service"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/user"
	matching_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/matching-service"
	user_service_usecase "github.com/Doremi203/couply/backend/matcher/internal/usecase/user-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		app.AddGatewayOptions(
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "x-api-key":
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

		dbClient, err := postgrespkg.NewClient(ctx, dbConfig)
		if err != nil {
			return errors.WrapFail(err, "create postgres client")
		}
		app.AddCloser(dbClient.Close)

		txManager := postgres.NewTxManager(dbClient)
		pgStorageUser := user.NewPgStorageUser(txManager)
		storageFacadeUser := user_service_facade.NewStorageFacadeUser(txManager, pgStorageUser)
		useCaseUserService := user_service_usecase.NewUseCase(storageFacadeUser)
		implUserService := user_service.NewImplementation(useCaseUserService)

		pgStorageMatching := matching.NewPgStorageMatching(txManager)
		storageFacadeMatching := matching_service_facade.NewStorageFacadeMatching(txManager, pgStorageMatching)
		useCaseMatchingService := matching_service_usecase.NewUseCase(storageFacadeMatching)
		implMatchingService := matching_service.NewImplementation(useCaseMatchingService)

		pgStorageSearch := search.NewPgStorageSearch(txManager)
		storageFacadeSearch := search_service_facade.NewStorageFacadeSearch(txManager, pgStorageSearch, pgStorageUser)
		useCaseSearchService := search_service_usecase.NewUseCase(storageFacadeSearch)
		implSearchService := search_service.NewImplementation(useCaseSearchService)

		app.RegisterGRPCServices(implUserService, implMatchingService, implSearchService)
		app.AddGatewayHandlers(implUserService, implMatchingService, implSearchService)

		return nil
	})
}
