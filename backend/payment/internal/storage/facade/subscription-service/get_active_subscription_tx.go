package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) GetActiveSubscriptionTx(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	var sub *subscription.Subscription
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err = f.storage.GetActiveSubscription(ctxTx, userID)
		if err != nil {
			return errors.WrapFail(err, "get active subscription")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get active subscription transaction")
	}

	return sub, nil
}
