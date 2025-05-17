package facade

import (
	"github.com/Doremi203/couply/backend/blocker/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/postgres/blocker"
)

type StorageFacadeBlocker struct {
	txManager postgres.TransactionManager
	storage   *blocker.PgStorageBlocker
}

func NewStorageFacadeBlocker(
	txManager postgres.TransactionManager,
	pgRepository *blocker.PgStorageBlocker,
) *StorageFacadeBlocker {
	return &StorageFacadeBlocker{
		txManager: txManager,
		storage:   pgRepository,
	}
}
