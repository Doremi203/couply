package payment

import "github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

type PgStoragePayment struct {
	txManager postgres.TransactionManager
}

func NewPgStoragePayment(txManager postgres.TransactionManager) *PgStoragePayment {
	return &PgStoragePayment{txManager: txManager}
}
