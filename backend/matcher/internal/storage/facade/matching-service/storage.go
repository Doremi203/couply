package matching_service

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/matching"
)

type StorageFacadeMatching struct {
	txManager postgres.TransactionManager
	storage   *matching.PgStorageMatching
}

func NewStorageFacadeMatching(
	txManager postgres.TransactionManager,
	pgRepository *matching.PgStorageMatching,
) *StorageFacadeMatching {
	return &StorageFacadeMatching{
		txManager: txManager,
		storage:   pgRepository,
	}
}
