package facade

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
)

type StorageFacadePayment struct {
	txManager storage.TransactionManager
	storage   *postgres2.PgStoragePayment
}

func NewStorageFacadePayment(
	txManager storage.TransactionManager,
	pgRepository *postgres2.PgStoragePayment,
) *StorageFacadePayment {
	return &StorageFacadePayment{
		txManager: txManager,
		storage:   pgRepository,
	}
}
