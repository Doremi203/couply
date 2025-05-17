package blocker

import "github.com/Doremi203/couply/backend/blocker/internal/storage/postgres"

type PgStorageBlocker struct {
	txManager postgres.TransactionManager
}

func NewPgStorageBlocker(txManager postgres.TransactionManager) *PgStorageBlocker {
	return &PgStorageBlocker{txManager: txManager}
}
