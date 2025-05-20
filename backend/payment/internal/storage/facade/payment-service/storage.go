package payment_service

import (
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres/payment"
)

type StorageFacadePayment struct {
	txManager postgres.TransactionManager
	storage   *payment.PgStoragePayment
}

func NewStorageFacadePayment(
	txManager postgres.TransactionManager,
	pgRepository *payment.PgStoragePayment,
) *StorageFacadePayment {
	return &StorageFacadePayment{
		txManager: txManager,
		storage:   pgRepository,
	}
}
