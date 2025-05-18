package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
)

type subscriptionStorageFacade interface {
	CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) (*subscription.Subscription, error)
}

type UseCase struct {
	subscriptionStorageFacade subscriptionStorageFacade
}

func NewUseCase(subscriptionStorageFacade subscriptionStorageFacade) *UseCase {
	return &UseCase{subscriptionStorageFacade: subscriptionStorageFacade}
}
