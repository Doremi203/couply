package matching

import "github.com/Doremi203/couply/backend/matcher/internal/storage/postgres"

type PgStorageMatching struct {
	txManager postgres.TransactionManager
}

func NewPgStorageMatching(txManager postgres.TransactionManager) *PgStorageMatching {
	return &PgStorageMatching{txManager: txManager}
}
