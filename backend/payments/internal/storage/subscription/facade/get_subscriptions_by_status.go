package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) GetSubscriptionsByStatusTx(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error) {
	subs, err := f.subscriptionStorage.GetSubscriptions(ctx, postgres.GetSubscriptionsOptions{
		SubscriptionStatus: status,
	})
	if err != nil {
		return nil, errors.WrapFail(err, "storage.GetSubscriptions")
	}

	return subs, nil
}
