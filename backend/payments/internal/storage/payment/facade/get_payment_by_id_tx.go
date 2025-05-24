package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	paymentpostgres "github.com/Doremi203/couply/backend/payments/internal/storage/payment/postgres"
	"github.com/google/uuid"
)

func (f *StorageFacadePayment) GetPaymentByIDTx(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error) {
	pay, err := f.paymentStorage.GetPayment(ctx, paymentpostgres.GetPaymentOptions{
		PaymentID: paymentID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage.GetPayment")
	}

	return pay, nil
}
