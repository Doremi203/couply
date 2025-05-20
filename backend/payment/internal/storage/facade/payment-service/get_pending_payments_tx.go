package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
)

func (f *StorageFacadePayment) GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error) {
	var p []*payment.Payment
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		p, err = f.storage.GetPendingPayments(ctxTx)
		if err != nil {
			return errors.WrapFail(err, "get pending payments")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "get pending payments transaction")
	}

	return p, nil
}
