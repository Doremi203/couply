package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) (*subscription.Subscription, error) {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.AddSubscription(ctxTx, newSubscription)
		if err != nil {
			return errors.WrapFail(err, "add subscription")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "create subscription transaction")
	}

	return newSubscription, nil
}
