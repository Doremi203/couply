package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
)

func (c *UseCase) CreatePayment(ctx context.Context, in *dto.CreatePaymentV1Request) (*dto.CreatePaymentV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	gatewayID, err := c.paymentGateway.CreatePayment(ctx, in.Amount, in.Currency)
	if err != nil {
		return nil, errors.Wrap(err, "paymentGateway.CreatePayment")
	}

	newPayment, err := dto.CreatePaymentRequestToPayment(in, userID, gatewayID)
	if err != nil {
		return nil, errors.Wrap(err, "dto.CreatePaymentRequestToPayment")
	}

	err = c.paymentStorageFacade.CreatePaymentTx(ctx, newPayment)
	if err != nil {
		return nil, errors.Wrap(err, "paymentStorageFacade.CreatePaymentTx")
	}

	c.startPaymentStatusUpdate(newPayment)

	return dto.PaymentToCreatePaymentResponse(newPayment), nil
}

func (c *UseCase) startPaymentStatusUpdate(payment *payment.Payment) {
	go c.updater.CheckAndUpdatePaymentStatusWithRetry(context.Background(), payment.ID, payment.GatewayID)
}
