package facade

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionDatesTx(ctx context.Context, subscriptionID uuid.UUID, startDate, endDate time.Time) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err := f.subscriptionStorage.GetSubscription(ctx, postgres.GetSubscriptionOptions{
			SubscriptionID: subscriptionID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetSubscription")
		}

		sub.StartDate = startDate
		sub.EndDate = endDate

		err = f.subscriptionStorage.UpdateSubscription(ctx, sub)
		if err != nil {
			return errors.Wrap(err, "storage.UpdateSubscription")
		}

		return nil
	})

	return err
}
