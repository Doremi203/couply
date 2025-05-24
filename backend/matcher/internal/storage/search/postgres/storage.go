package postgres

import (
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
)

type PgStorageSearch struct {
	txManager storage.TransactionManager
}

func NewPgStorageSearch(txManager storage.TransactionManager) *PgStorageSearch {
	return &PgStorageSearch{txManager: txManager}
}
