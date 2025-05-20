package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) GetSubscriptionsByStatusTx(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error) {
	var subs []*subscription.Subscription
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		subs, err = f.subscriptionStorage.GetSubscriptionsByStatus(ctxTx, status)
		if err != nil {
			return errors.WrapFail(err, "get subscriptions by status")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get subscriptions by status transaction")
	}

	return subs, nil
}
