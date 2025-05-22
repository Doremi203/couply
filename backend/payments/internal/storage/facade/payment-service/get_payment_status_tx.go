package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) GetPaymentStatusTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error) {
	pay, err := f.storage.GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, errors.Wrap(err, "GetPaymentStatusTx")
	}

	return pay, nil
}
