package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
)

type subscriptionStorageFacade interface {
	CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) (*subscription.Subscription, error)
	GetActiveSubscriptionTx(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error)
	CancelSubscriptionTx(ctx context.Context, subscriptionID uuid.UUID) error
}

type UseCase struct {
	subscriptionStorageFacade subscriptionStorageFacade
}

func NewUseCase(subscriptionStorageFacade subscriptionStorageFacade) *UseCase {
	return &UseCase{subscriptionStorageFacade: subscriptionStorageFacade}
}
