package payment_service

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/payment-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreatePaymentV1Request struct {
	SubscriptionID uuid.UUID
	Amount         int64
	Currency       string
}

func PBToCreatePaymentRequest(req *desc.CreatePaymentV1Request) (*CreatePaymentV1Request, error) {
	subID, err := uuid.Parse(req.GetSubscriptionId())
	if err != nil {
		return nil, err
	}
	return &CreatePaymentV1Request{
		SubscriptionID: subID,
		Amount:         req.GetAmount(),
		Currency:       req.GetCurrency(),
	}, nil
}

func CreatePaymentRequestToPayment(req *CreatePaymentV1Request, userID uuid.UUID, gatewayID string) (*payment.Payment, error) {
	paymentID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "CreatePaymentRequestToPayment")
	}

	now := time.Now()
	return &payment.Payment{
		ID:             paymentID,
		UserID:         userID,
		SubscriptionID: req.SubscriptionID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         payment.PaymentStatusPending,
		GatewayID:      gatewayID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

type CreatePaymentV1Response struct {
	PaymentID string
	Status    payment.PaymentStatus
	UpdatedAt time.Time
}

func CreatePaymentResponseToPB(resp *CreatePaymentV1Response) *desc.CreatePaymentV1Response {
	return &desc.CreatePaymentV1Response{
		PaymentId: resp.PaymentID,
		Status:    payment.PaymentStatusToPB(resp.Status),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

func PaymentToCreatePaymentResponse(pay *payment.Payment) *CreatePaymentV1Response {
	return &CreatePaymentV1Response{
		PaymentID: pay.ID.String(),
		Status:    pay.Status,
		UpdatedAt: pay.UpdatedAt,
	}
}
