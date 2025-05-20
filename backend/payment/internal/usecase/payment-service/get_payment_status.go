package payment_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/payment/internal/dto/payment-service"
)

func (c *UseCase) GetPaymentStatus(ctx context.Context, in *dto.GetPaymentStatusV1Request) (*dto.GetPaymentStatusV1Response, error) {
	payment, err := c.paymentStorageFacade.GetPaymentStatusTx(ctx, in.GetPaymentID())
	if err != nil {
		return nil, err
	}

	return &dto.GetPaymentStatusV1Response{
		PaymentID:     payment.GetID(),
		PaymentStatus: payment.GetStatus(),
		UpdatedAt:     payment.GetUpdatedAt(),
	}, nil
}
