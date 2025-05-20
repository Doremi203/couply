package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) CancelSubscriptionTx(ctx context.Context, subscriptionID uuid.UUID) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.subscriptionStorage.UpdateSubscription(ctxTx, subscriptionID, subscription.SubscriptionStatusCanceled)
		if err != nil {
			return errors.WrapFail(err, "update subscription")
		}

		return nil
	})

	if err != nil {
		return errors.WrapFail(err, "update subscription transaction")
	}

	return nil
}
