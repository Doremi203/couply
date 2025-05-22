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
		sub, err = f.getSubscriptionWithPayments(ctxTx, userID)
		return err
	})

	if err != nil {
		return nil, errors.Wrap(err, "GetActiveSubscriptionTx")
	}

	return sub, nil
}

func (f *StorageFacadeSubscription) getSubscriptionWithPayments(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	sub, err := f.subscriptionStorage.GetActiveSubscriptionByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	paymentIDs, err := f.getSubscriptionPaymentIDs(ctx, sub.GetID())
	if err != nil {
		return nil, err
	}

	sub.PaymentIDs = paymentIDs
	return sub, nil
}

func (f *StorageFacadeSubscription) getSubscriptionPaymentIDs(ctx context.Context, subscriptionID uuid.UUID) ([]uuid.UUID, error) {
	ids, err := f.paymentStorage.GetSubscriptionPayments(ctx, subscriptionID)
	if err != nil {
		return nil, err
	}

	return convertStringsToUUIDs(ids)
}

func convertStringsToUUIDs(ids []string) ([]uuid.UUID, error) {
	uuids := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		parsedUUID, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		uuids[i] = parsedUUID
	}
	return uuids, nil
}
