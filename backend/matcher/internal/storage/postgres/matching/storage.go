package matching

import "github.com/Doremi203/Couply/backend/internal/storage/postgres"

type PgStorageMatching struct {
	txManager postgres.TransactionManager
}

func NewPgStorageMatching(txManager postgres.TransactionManager) *PgStorageMatching {
	return &PgStorageMatching{txManager: txManager}
}
