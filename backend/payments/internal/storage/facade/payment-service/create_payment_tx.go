package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
)

func (f *StorageFacadePayment) CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) (*payment.Payment, error) {
	if err := f.storage.AddPayment(ctx, newPayment); err != nil {
		return nil, errors.Wrap(err, "CreatePaymentTx")
	}

	return newPayment, nil
}
