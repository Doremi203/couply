package matching_service

import (
	"github.com/Doremi203/Couply/backend/internal/storage/postgres"
	"github.com/Doremi203/Couply/backend/internal/storage/postgres/matching"
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
