package payment_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/payment-service/v1"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreatePaymentV1(ctx context.Context, in *desc.CreatePaymentV1Request) (*desc.CreatePaymentV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToCreatePaymentRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.CreatePayment(ctx, req)
	switch {
	case errors.Is(err, payment.ErrSubscriptionDoesntExist):
		return nil, status.Error(codes.FailedPrecondition, payment.ErrSubscriptionDoesntExist.Error())
	case errors.Is(err, payment.ErrDuplicatePayment):
		return nil, status.Error(codes.FailedPrecondition, payment.ErrDuplicatePayment.Error())
	case err != nil:
		return nil, err
	}

	return dto.CreatePaymentResponseToPB(response), nil
}
