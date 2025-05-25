package postgres

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage"
)

type PgStoragePayment struct {
	txManager storage.TransactionManager
}

func NewPgStoragePayment(txManager storage.TransactionManager) *PgStoragePayment {
	return &PgStoragePayment{txManager: txManager}
}
