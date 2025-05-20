package subscription_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionDatesTx(ctx context.Context, subscriptionID uuid.UUID, startDate, endDate time.Time) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.subscriptionStorage.UpdateSubscriptionDates(ctxTx, subscriptionID, startDate, endDate)
		if err != nil {
			return errors.WrapFail(err, "update subscription dates")
		}

		return nil
	})

	if err != nil {
		return errors.WrapFail(err, "update subscription dates transaction")
	}

	return nil
}
