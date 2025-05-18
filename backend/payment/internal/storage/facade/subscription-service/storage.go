package subscription_service

import (
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres"
	"github.com/Doremi203/couply/backend/payment/internal/storage/postgres/subscription"
)

type StorageFacadeSubscription struct {
	txManager postgres.TransactionManager
	storage   *subscription.PgStorageSubscription
}

func NewStorageFacadeSubscription(
	txManager postgres.TransactionManager,
	pgRepository *subscription.PgStorageSubscription,
) *StorageFacadeSubscription {
	return &StorageFacadeSubscription{
		txManager: txManager,
		storage:   pgRepository,
	}
}
