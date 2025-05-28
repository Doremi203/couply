package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
)

type subscriptionServiceStorage interface {
	GetSubscription(ctx context.Context, opts postgres.GetSubscriptionOptions) (*subscription.Subscription, error)
}

type paymentServiceStorage interface {
	paymentStorage
}

type paymentStorage interface {
	CreatePayment(ctx context.Context, payment *payment.Payment) error
	GetPayment(ctx context.Context, opts postgres2.GetPaymentOptions) (*payment.Payment, error)
	GetPayments(ctx context.Context, opts postgres2.GetPaymentsOptions) ([]*payment.Payment, error)
	UpdatePayment(ctx context.Context, pay *payment.Payment) error
}

type StorageFacadePayment struct {
	txManager           storage.TransactionManager
	paymentStorage      paymentServiceStorage
	subscriptionStorage subscriptionServiceStorage
}

func NewStorageFacadePayment(
	txManager storage.TransactionManager,
	pgPaymentRepository paymentServiceStorage,
	pgSubscriptionRepository subscriptionServiceStorage,
) *StorageFacadePayment {
	return &StorageFacadePayment{
		txManager:           txManager,
		paymentStorage:      pgPaymentRepository,
		subscriptionStorage: pgSubscriptionRepository,
	}
}
