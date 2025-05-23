package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

type subscriptionStorageFacade interface {
	subscriptionStorageSetterFacade
	subscriptionStorageGetterFacade
}

type subscriptionStorageSetterFacade interface {
	CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) error
	UpdateSubscriptionStatusTx(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error
}

type subscriptionStorageGetterFacade interface {
	GetActiveSubscriptionTx(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error)
}

type UseCase struct {
	subscriptionStorageFacade subscriptionStorageFacade
}

func NewUseCase(subscriptionStorageFacade subscriptionStorageFacade) *UseCase {
	return &UseCase{subscriptionStorageFacade: subscriptionStorageFacade}
}
