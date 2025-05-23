package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	payment2 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		pay, err := f.storage.GetPayment(ctxTx, payment2.GetPaymentOptions{
			PaymentID: paymentID,
			ForUpdate: true,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetPayment")
		}

		pay.Status = newStatus

		err = f.storage.UpdatePayment(ctx, pay)
		if err != nil {
			return errors.Wrap(err, "storage.UpdatePayment")
		}

		return nil
	})

	return err
}
