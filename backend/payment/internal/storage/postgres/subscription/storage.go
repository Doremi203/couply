package subscription

import "github.com/Doremi203/couply/backend/payment/internal/storage/postgres"

type PgStorageSubscription struct {
	txManager postgres.TransactionManager
}

func NewPgStorageSubscription(txManager postgres.TransactionManager) *PgStorageSubscription {
	return &PgStorageSubscription{txManager: txManager}
}
