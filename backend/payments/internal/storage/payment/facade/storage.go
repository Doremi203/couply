package facade

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
)

type StorageFacadePayment struct {
	txManager           storage.TransactionManager
	paymentStorage      *postgres2.PgStoragePayment
	subscriptionStorage *postgres.PgStorageSubscription
}

func NewStorageFacadePayment(
	txManager storage.TransactionManager,
	pgPaymentRepository *postgres2.PgStoragePayment,
	pgSubscriptionRepository *postgres.PgStorageSubscription,
) *StorageFacadePayment {
	return &StorageFacadePayment{
		txManager:           txManager,
		paymentStorage:      pgPaymentRepository,
		subscriptionStorage: pgSubscriptionRepository,
	}
}
