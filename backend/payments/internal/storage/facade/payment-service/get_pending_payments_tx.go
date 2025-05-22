package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
)

func (f *StorageFacadePayment) GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error) {
	pays, err := f.storage.GetPendingPayments(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetPendingPaymentsTx")
	}

	return pays, nil
}
