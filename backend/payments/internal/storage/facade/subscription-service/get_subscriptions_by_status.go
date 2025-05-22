package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) GetSubscriptionsByStatusTx(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error) {
	subs, err := f.subscriptionStorage.GetSubscriptionsByStatus(ctx, status)
	if err != nil {
		return nil, errors.WrapFail(err, "GetSubscriptionsByStatusTx")
	}

	return subs, nil
}
