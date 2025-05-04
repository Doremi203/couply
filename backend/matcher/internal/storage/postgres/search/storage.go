package search

import "github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"

type PgStorageSearch struct {
	txManager postgres.TransactionManager
}

func NewPgStorageSearch(txManager postgres.TransactionManager) *PgStorageSearch {
	return &PgStorageSearch{txManager: txManager}
}
