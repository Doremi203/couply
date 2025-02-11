package facade

import "github.com/Doremi203/Couply/backend/internal/storage/postgres"

type StorageFacade struct {
	txManager postgres.TransactionManager
	storage   *postgres.PgStorage
}

func NewStorageFacade(
	txManager postgres.TransactionManager,
	pgRepository *postgres.PgStorage,
) *StorageFacade {
	return &StorageFacade{
		txManager: txManager,
		storage:   pgRepository,
	}
}
