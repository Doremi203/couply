package postgres

import (
	"github.com/Doremi203/couply/backend/blocker/internal/storage"
)

type PgStorageBlocker struct {
	txManager storage.TransactionManager
}

func NewPgStorageBlocker(txManager storage.TransactionManager) *PgStorageBlocker {
	return &PgStorageBlocker{txManager: txManager}
}
