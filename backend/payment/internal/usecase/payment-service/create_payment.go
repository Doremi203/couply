package payment_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/payment-service"
	"github.com/Doremi203/couply/backend/payment/utils"
	"github.com/google/uuid"
)

func (c *UseCase) CreatePayment(ctx context.Context, in *dto.CreatePaymentV1Request) (*dto.CreatePaymentV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	paymentID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	gatewayID, err := c.paymentGateway.CreatePayment(ctx, in.GetAmount(), in.GetCurrency())
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newPayment := &payment.Payment{
		ID:             paymentID,
		UserID:         userID,
		SubscriptionID: in.GetSubscriptionID(),
		Amount:         in.GetAmount(),
		Currency:       in.GetCurrency(),
		Status:         payment.PaymentStatusPending,
		GatewayID:      gatewayID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	createdPayment, err := c.paymentStorageFacade.CreatePaymentTx(ctx, newPayment)
	if err != nil {
		return nil, err
	}

	go c.updater.CheckAndUpdatePaymentStatusWithRetry(context.Background(), createdPayment.GetID(), createdPayment.GetGatewayID())

	return &dto.CreatePaymentV1Response{
		PaymentID: createdPayment.GetID().String(),
		Status:    createdPayment.GetStatus(),
		UpdatedAt: createdPayment.GetUpdatedAt(),
	}, nil
}
