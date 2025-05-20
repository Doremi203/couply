package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error) {
	var p *payment.Payment
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		p, err = f.storage.GetPayment(ctxTx, paymentID)
		if err != nil {
			return errors.WrapFail(err, "get payment")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get payment status transaction")
	}

	return p, nil
}
