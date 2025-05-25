package postgres

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
)

type PgStorageMatching struct {
	txManager storage.TransactionManager
}

func NewPgStorageMatching(txManager storage.TransactionManager) *PgStorageMatching {
	return &PgStorageMatching{txManager: txManager}
}
