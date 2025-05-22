package payment_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"time"

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

func (x *CreatePaymentV1Request) GetSubscriptionID() uuid.UUID {
	if x != nil {
		return x.SubscriptionID
	}
	return uuid.Nil
}

func (x *CreatePaymentV1Request) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreatePaymentV1Request) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
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
		SubscriptionID: req.GetSubscriptionID(),
		Amount:         req.GetAmount(),
		Currency:       req.GetCurrency(),
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

func (x *CreatePaymentV1Response) GetPaymentID() string {
	if x != nil {
		return x.PaymentID
	}
	return ""
}

func (x *CreatePaymentV1Response) GetStatus() payment.PaymentStatus {
	if x != nil {
		return x.Status
	}
	return payment.PaymentStatus(0)
}

func (x *CreatePaymentV1Response) GetUpdatedAt() time.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return time.Time{}
}

func CreatePaymentResponseToPB(resp *CreatePaymentV1Response) *desc.CreatePaymentV1Response {
	return &desc.CreatePaymentV1Response{
		PaymentId: resp.GetPaymentID(),
		Status:    payment.PaymentStatusToPB(resp.GetStatus()),
		UpdatedAt: timestamppb.New(resp.GetUpdatedAt()),
	}
}

func PaymentToCreatePaymentResponse(pay *payment.Payment) *CreatePaymentV1Response {
	return &CreatePaymentV1Response{
		PaymentID: pay.GetID().String(),
		Status:    pay.GetStatus(),
		UpdatedAt: pay.GetUpdatedAt(),
	}
}
