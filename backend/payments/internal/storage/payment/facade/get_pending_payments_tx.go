package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	payment2 "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
)

func (f *StorageFacadePayment) GetPendingPaymentsTx(ctx context.Context) ([]*payment.Payment, error) {
	pays, err := f.paymentStorage.GetPayments(ctx, payment2.GetPaymentsOptions{
		PendingPayments: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage.GetPayments")
	}

	return pays, nil
}
