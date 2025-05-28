package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	postgres3 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
)

type paymentServiceStorage interface {
	GetPayments(ctx context.Context, opts postgres3.GetPaymentsOptions) ([]*payment.Payment, error)
}

type subscriptionServiceStorage interface {
	subscriptionStorage
}

type subscriptionStorage interface {
	CreateSubscription(ctx context.Context, subscription *subscription.Subscription) error
	GetSubscription(ctx context.Context, opts postgres2.GetSubscriptionOptions) (*subscription.Subscription, error)
	GetSubscriptions(ctx context.Context, opts postgres2.GetSubscriptionsOptions) ([]*subscription.Subscription, error)
	UpdateSubscription(ctx context.Context, sub *subscription.Subscription) error
}

type StorageFacadeSubscription struct {
	txManager           storage.TransactionManager
	subscriptionStorage subscriptionServiceStorage
	paymentStorage      paymentServiceStorage
}

func NewStorageFacadeSubscription(
	txManager storage.TransactionManager,
	pgSubscriptionRepository subscriptionServiceStorage,
	pgPaymentRepository paymentServiceStorage,
) *StorageFacadeSubscription {
	return &StorageFacadeSubscription{
		txManager:           txManager,
		subscriptionStorage: pgSubscriptionRepository,
		paymentStorage:      pgPaymentRepository,
	}
}
