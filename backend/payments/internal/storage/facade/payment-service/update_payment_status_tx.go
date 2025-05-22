package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) UpdatePaymentStatusTx(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error {
	err := f.storage.UpdatePaymentStatus(ctx, paymentID, newStatus)
	if err != nil {
		return errors.Wrap(err, "UpdatePaymentStatusTx")
	}

	return nil
}
