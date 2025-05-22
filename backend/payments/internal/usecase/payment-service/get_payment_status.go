package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
)

func (c *UseCase) GetPaymentStatus(ctx context.Context, in *dto.GetPaymentStatusV1Request) (*dto.GetPaymentStatusV1Response, error) {
	payment, err := c.paymentStorageFacade.GetPaymentStatusTx(ctx, in.PaymentID)
	if err != nil {
		return nil, errors.Wrap(err, "GetPaymentStatus")
	}

	return dto.PaymentToGetPaymentStatusResponse(payment), nil
}
