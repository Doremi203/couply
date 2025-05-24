package facade

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"
)

type StorageFacadeUser struct {
	txManager storage.TransactionManager
	storage   *postgres2.PgStorageUser
}

func NewStorageFacadeUser(
	txManager storage.TransactionManager,
	pgRepository *postgres2.PgStorageUser,
) *StorageFacadeUser {
	return &StorageFacadeUser{
		txManager: txManager,
		storage:   pgRepository,
	}
}
