package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionStatusTx(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.subscriptionStorage.UpdateSubscriptionStatus(ctxTx, subscriptionID, status)
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
