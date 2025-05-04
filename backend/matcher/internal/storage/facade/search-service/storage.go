package search_service

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/search"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/user"
)

type StorageFacadeSearch struct {
	txManager     postgres.TransactionManager
	searchStorage *search.PgStorageSearch
	userStorage   *user.PgStorageUser
}

func NewStorageFacadeSearch(
	txManager postgres.TransactionManager,
	pgRepositorySearch *search.PgStorageSearch,
	pgRepositoryUser *user.PgStorageUser,
) *StorageFacadeSearch {
	return &StorageFacadeSearch{
		txManager:     txManager,
		searchStorage: pgRepositorySearch,
		userStorage:   pgRepositoryUser,
	}
}
