package user_service

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/user"
)

type StorageFacadeUser struct {
	txManager postgres.TransactionManager
	storage   *user.PgStorageUser
}

func NewStorageFacadeUser(
	txManager postgres.TransactionManager,
	pgRepository *user.PgStorageUser,
) *StorageFacadeUser {
	return &StorageFacadeUser{
		txManager: txManager,
		storage:   pgRepository,
	}
}
