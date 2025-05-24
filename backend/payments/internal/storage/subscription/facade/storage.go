package facade

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	postgres3 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
)

type StorageFacadeSubscription struct {
	txManager           storage.TransactionManager
	subscriptionStorage *postgres2.PgStorageSubscription
	paymentStorage      *postgres3.PgStoragePayment
}

func NewStorageFacadeSubscription(
	txManager storage.TransactionManager,
	pgSubscriptionRepository *postgres2.PgStorageSubscription,
	pgPaymentRepository *postgres3.PgStoragePayment,
) *StorageFacadeSubscription {
	return &StorageFacadeSubscription{
		txManager:           txManager,
		subscriptionStorage: pgSubscriptionRepository,
		paymentStorage:      pgPaymentRepository,
	}
}
