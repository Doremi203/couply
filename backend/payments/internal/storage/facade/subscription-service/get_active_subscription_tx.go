package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) GetActiveSubscriptionTx(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	var sub *subscription.Subscription
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err = f.subscriptionStorage.GetActiveSubscription(ctxTx, userID)
		if err != nil {
			return errors.WrapFail(err, "get active subscription")
		}

		ids, err := f.paymentStorage.GetPaymentIDsBySubscriptionID(ctxTx, sub.GetID())
		if err != nil {
			return errors.WrapFail(err, "get active subscription ids")
		}

		uuids := make([]uuid.UUID, len(ids))
		for i, id := range ids {
			parsedUUID, err := uuid.Parse(id)
			if err != nil {
				return errors.WrapFail(err, "parse active subscription uuid")
			}
			uuids[i] = parsedUUID
		}

		sub.PaymentIDs = uuids

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get active subscription transaction")
	}

	return sub, nil
}
