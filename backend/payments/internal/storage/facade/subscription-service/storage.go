package subscription_service

import (
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres/payment"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres/subscription"
)

type StorageFacadeSubscription struct {
	txManager           postgres.TransactionManager
	subscriptionStorage *subscription.PgStorageSubscription
	paymentStorage      *payment.PgStoragePayment
}

func NewStorageFacadeSubscription(
	txManager postgres.TransactionManager,
	pgSubscriptionRepository *subscription.PgStorageSubscription,
	pgPaymentRepository *payment.PgStoragePayment,
) *StorageFacadeSubscription {
	return &StorageFacadeSubscription{
		txManager:           txManager,
		subscriptionStorage: pgSubscriptionRepository,
		paymentStorage:      pgPaymentRepository,
	}
}
