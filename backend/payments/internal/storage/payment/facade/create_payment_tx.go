package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
)

func (f *StorageFacadePayment) CreatePaymentTx(ctx context.Context, newPayment *payment.Payment) error {
	if err := f.storage.CreatePayment(ctx, newPayment); err != nil {
		return errors.Wrap(err, "storage.CreatePayment")
	}

	return nil
}
