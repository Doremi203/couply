package postgres

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
)

type PgStorageUser struct {
	txManager storage.TransactionManager
}

func NewPgStorageUser(txManager storage.TransactionManager) *PgStorageUser {
	return &PgStorageUser{txManager: txManager}
}
