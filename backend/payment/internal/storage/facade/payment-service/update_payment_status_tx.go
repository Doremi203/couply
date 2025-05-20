package payment_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.UpdatePaymentStatus(ctxTx, paymentID, newStatus)
		if err != nil {
			return errors.WrapFail(err, "update payment status")
		}

		return nil
	})

	if err != nil {
		return errors.WrapFail(err, "update payment status transaction")
	}

	return nil
}
