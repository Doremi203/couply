package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
)

func (f *StorageFacadePayment) CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error) {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.AddPayment(ctxTx, newPayment)
		if err != nil {
			return errors.WrapFail(err, "add payment")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "create payment transaction")
	}

	return newPayment, nil
}
