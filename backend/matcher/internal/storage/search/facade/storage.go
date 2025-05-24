package facade

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres3 "github.com/Doremi203/couply/backend/matcher/internal/storage/search/postgres"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"
)

type StorageFacadeSearch struct {
	txManager     storage.TransactionManager
	searchStorage *postgres3.PgStorageSearch
	userStorage   *postgres2.PgStorageUser
}

func NewStorageFacadeSearch(
	txManager storage.TransactionManager,
	pgRepositorySearch *postgres3.PgStorageSearch,
	pgRepositoryUser *postgres2.PgStorageUser,
) *StorageFacadeSearch {
	return &StorageFacadeSearch{
		txManager:     txManager,
		searchStorage: pgRepositorySearch,
		userStorage:   pgRepositoryUser,
	}
}
