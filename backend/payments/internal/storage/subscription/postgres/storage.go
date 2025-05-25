package postgres

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage"
)

type PgStorageSubscription struct {
	txManager storage.TransactionManager
}

func NewPgStorageSubscription(txManager storage.TransactionManager) *PgStorageSubscription {
	return &PgStorageSubscription{txManager: txManager}
}
