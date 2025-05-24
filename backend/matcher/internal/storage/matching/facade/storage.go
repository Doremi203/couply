package facade

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"
)

type StorageFacadeMatching struct {
	txManager storage.TransactionManager
	storage   *postgres2.PgStorageMatching
}

func NewStorageFacadeMatching(
	txManager storage.TransactionManager,
	pgRepository *postgres2.PgStorageMatching,
) *StorageFacadeMatching {
	return &StorageFacadeMatching{
		txManager: txManager,
		storage:   pgRepository,
	}
}
