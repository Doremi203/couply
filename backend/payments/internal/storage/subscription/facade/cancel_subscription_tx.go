package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) CancelSubscriptionTx(ctx context.Context, subscriptionID uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err := f.subscriptionStorage.GetSubscription(ctxTx, postgres.GetSubscriptionOptions{
			SubscriptionID: subscriptionID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetSubscription")
		}

		if sub.Status != subscription.SubscriptionStatusActive {
			return subscription.ErrSubscriptionIsNotActive
		}

		sub.Status = subscription.SubscriptionStatusCanceled

		err = f.subscriptionStorage.UpdateSubscription(ctxTx, sub)
		if err != nil {
			return errors.Wrap(err, "storage.UpdateSubscription")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
