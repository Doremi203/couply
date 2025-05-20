package subscription

import "github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

type PgStorageSubscription struct {
	txManager postgres.TransactionManager
}

func NewPgStorageSubscription(txManager postgres.TransactionManager) *PgStorageSubscription {
	return &PgStorageSubscription{txManager: txManager}
}
