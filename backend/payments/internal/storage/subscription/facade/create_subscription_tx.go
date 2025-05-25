package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if ok := f.subscriptionExists(ctxTx, newSubscription); ok {
			return errors.Wrap(subscription.ErrAlreadyActiveSubscriptionExists, "checkExistingSubscription")
		}

		err := f.subscriptionStorage.CreateSubscription(ctxTx, newSubscription)
		if err != nil {
			return errors.Wrap(err, "storage.CreateSubscription")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}

func (f *StorageFacadeSubscription) subscriptionExists(ctx context.Context, newSubscription *subscription.Subscription) bool {
	_, err := f.subscriptionStorage.GetSubscription(ctx, postgres.GetSubscriptionOptions{
		UserID:                      newSubscription.UserID,
		ActiveOrPendingSubscription: true,
	})
	if err != nil {
		return false
	}

	return true
}
