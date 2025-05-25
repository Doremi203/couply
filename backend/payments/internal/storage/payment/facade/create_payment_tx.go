package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage/subscription/postgres"
)

func (f *StorageFacadePayment) CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		sub, err := f.subscriptionStorage.GetSubscription(ctxTx, postgres.GetSubscriptionOptions{
			SubscriptionID: newPayment.SubscriptionID,
		})
		if err != nil {
			return errors.Wrap(err, "subscriptionStorage.GetSubscription")
		}

		if sub.Status == subscription.SubscriptionStatusActive {
			return subscription.ErrSubscriptionHasAlreadyBeenPaid
		}

		if err = f.paymentStorage.CreatePayment(ctxTx, newPayment); err != nil {
			return errors.Wrap(err, "storage.CreatePayment")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
