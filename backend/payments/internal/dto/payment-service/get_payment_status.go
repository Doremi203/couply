package payment_service

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/payment-service/v1"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetPaymentStatusV1Request struct {
	PaymentID uuid.UUID
}

type GetPaymentStatusV1Response struct {
	PaymentID     uuid.UUID
	PaymentStatus payment.PaymentStatus
	UpdatedAt     time.Time
}

func PBToGetPaymentStatusRequest(req *desc.GetPaymentStatusV1Request) (*GetPaymentStatusV1Request, error) {
	paymentID, err := uuid.Parse(req.GetPaymentId())
	if err != nil {
		return nil, errors.Wrap(err, "PBToCreatePaymentRequest")
	}
	return &GetPaymentStatusV1Request{
		PaymentID: paymentID,
	}, nil
}

func GetPaymentStatusResponseToPB(resp *GetPaymentStatusV1Response) *desc.GetPaymentStatusV1Response {
	return &desc.GetPaymentStatusV1Response{
		PaymentId: resp.PaymentID.String(),
		Status:    payment.PaymentStatusToPB(resp.PaymentStatus),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

func PaymentToGetPaymentStatusResponse(pay *payment.Payment) *GetPaymentStatusV1Response {
	return &GetPaymentStatusV1Response{
		PaymentID:     pay.ID,
		PaymentStatus: pay.Status,
		UpdatedAt:     pay.UpdatedAt,
	}
}
