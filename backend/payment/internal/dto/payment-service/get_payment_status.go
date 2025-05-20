package payment_service

import (
	"time"

	desc "github.com/Doremi203/couply/backend/payment/gen/api/payment-service/v1"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetPaymentStatusV1Request struct {
	PaymentID uuid.UUID
}

func (x *GetPaymentStatusV1Request) GetPaymentID() uuid.UUID {
	if x != nil {
		return x.PaymentID
	}
	return uuid.Nil
}

type GetPaymentStatusV1Response struct {
	PaymentID     uuid.UUID
	PaymentStatus payment.PaymentStatus
	UpdatedAt     time.Time
}

func (x *GetPaymentStatusV1Response) GetPaymentID() uuid.UUID {
	if x != nil {
		return x.PaymentID
	}
	return uuid.Nil
}

func (x *GetPaymentStatusV1Response) GetPaymentStatus() payment.PaymentStatus {
	if x != nil {
		return x.PaymentStatus
	}
	return payment.PaymentStatus(0)
}

func (x *GetPaymentStatusV1Response) GetUpdatedAt() time.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return time.Time{}
}

func PBToGetPaymentStatusRequest(req *desc.GetPaymentStatusV1Request) (*GetPaymentStatusV1Request, error) {
	paymentID, err := uuid.Parse(req.GetPaymentId())
	if err != nil {
		return nil, err
	}
	return &GetPaymentStatusV1Request{
		PaymentID: paymentID,
	}, nil
}

func GetPaymentStatusResponseToPB(resp *GetPaymentStatusV1Response) *desc.GetPaymentStatusV1Response {
	return &desc.GetPaymentStatusV1Response{
		PaymentId: resp.GetPaymentID().String(),
		Status:    payment.PaymentStatusToPB(resp.GetPaymentStatus()),
		UpdatedAt: timestamppb.New(resp.GetUpdatedAt()),
	}
}
