package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) (*subscription.Subscription, error) {
	err := f.subscriptionStorage.AddSubscription(ctx, newSubscription)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSubscriptionTx")
	}

	return newSubscription, nil
}
