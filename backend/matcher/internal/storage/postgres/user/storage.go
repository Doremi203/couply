package user

import "github.com/Doremi203/Couply/backend/internal/storage/postgres"

type PgStorageUser struct {
	txManager postgres.TransactionManager
}

func NewPgStorageUser(txManager postgres.TransactionManager) *PgStorageUser {
	return &PgStorageUser{txManager: txManager}
}
