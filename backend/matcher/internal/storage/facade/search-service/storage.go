package search_service

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/search"
)

type StorageFacadeSearch struct {
	txManager postgres.TransactionManager
	storage   *search.PgStorageSearch
}

func NewStorageFacadeSearch(
	txManager postgres.TransactionManager,
	pgRepository *search.PgStorageSearch,
) *StorageFacadeSearch {
	return &StorageFacadeSearch{
		txManager: txManager,
		storage:   pgRepository,
	}
}
