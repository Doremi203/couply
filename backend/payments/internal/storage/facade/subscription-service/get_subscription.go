package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) GetSubscriptionTx(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error) {
	var sub *subscription.Subscription
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err = f.subscriptionStorage.GetSubscriptionByID(ctxTx, subID)
		if err != nil {
			return errors.WrapFail(err, "get subscription")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get subscription transaction")
	}

	return sub, nil
}
