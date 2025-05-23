package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	postgres2 "github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) GetActiveSubscriptionTx(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	var sub *subscription.Subscription
	var pays []*payment.Payment
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err = f.subscriptionStorage.GetSubscription(ctxTx, postgres2.GetSubscriptionOptions{
			UserID:             userID,
			ActiveSubscription: true,
		})
		if err != nil {
			if errors.Is(err, subscription.ErrSubscriptionNotFound) {
				return subscription.ErrActiveSubscriptionDoesntExist
			}
			return errors.Wrap(err, "storage.GetSubscription")
		}

		// getting payment ids cuz result will be returned to user
		pays, err = f.paymentStorage.GetPayments(ctxTx, postgres.GetPaymentsOptions{
			SubscriptionID: sub.ID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetPayments")
		}

		sub.PaymentIDs = getIDs(pays)

		return nil
	})

	return sub, err
}

func getIDs(pays []*payment.Payment) []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(pays))
	for _, pay := range pays {
		ids = append(ids, pay.ID)
	}

	return ids
}
