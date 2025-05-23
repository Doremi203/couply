package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionStatusTx(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err := f.subscriptionStorage.GetSubscription(ctx, postgres.GetSubscriptionOptions{
			SubscriptionID: subscriptionID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetSubscription")
		}

		sub.Status = status

		err = f.subscriptionStorage.UpdateSubscription(ctx, sub)
		if err != nil {
			return errors.Wrap(err, "storage.UpdateSubscription")
		}

		return nil
	})

	return err
}
