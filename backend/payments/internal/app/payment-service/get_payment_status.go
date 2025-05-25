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

func (i *Implementation) GetPaymentStatusV1(ctx context.Context, in *desc.GetPaymentStatusV1Request) (*desc.GetPaymentStatusV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToGetPaymentStatusRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetPaymentStatus(ctx, req)
	switch {
	case errors.Is(err, payment.ErrPaymentNotFound):
		return nil, status.Error(codes.NotFound, payment.ErrPaymentNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.GetPaymentStatusResponseToPB(response), nil
}
