package facade

import (
	"github.com/Doremi203/couply/backend/blocker/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"
)

type StorageFacadeBlocker struct {
	txManager storage.TransactionManager
	storage   *postgres2.PgStorageBlocker
}

func NewStorageFacadeBlocker(
	txManager storage.TransactionManager,
	pgRepository *postgres2.PgStorageBlocker,
) *StorageFacadeBlocker {
	return &StorageFacadeBlocker{
		txManager: txManager,
		storage:   pgRepository,
	}
}
