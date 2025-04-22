package user

import "github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"

type PgStorageUser struct {
	txManager postgres.TransactionManager
}

func NewPgStorageUser(txManager postgres.TransactionManager) *PgStorageUser {
	return &PgStorageUser{txManager: txManager}
}
