package facade

import (
	"context"
	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err := f.checkExistingSubscription(ctx, newSubscription)
		if err != nil {
			return errors.Wrap(err, "checkExistingSubscription")
		}

		err = f.subscriptionStorage.CreateSubscription(ctx, newSubscription)
		if err != nil {
			return errors.Wrap(err, "storage.CreateSubscription")
		}

		return nil
	})

	return err
}

func (f *StorageFacadeSubscription) checkExistingSubscription(ctx context.Context, newSubscription *subscription.Subscription) error {
	_, err := f.subscriptionStorage.GetSubscription(ctx, postgres.GetSubscriptionOptions{
		UserID:             newSubscription.UserID,
		ActiveSubscription: true,
	})
	if err != nil {
		if errors.Is(err, subscription.ErrSubscriptionNotFound) {
			return errors.Wrap(err, "subscriptionStorage.GetSubscription")
		}
		return subscription.ErrAlreadyActiveSubscriptionExists
	}

	return nil
}
