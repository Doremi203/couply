package payment_service

import (
	"context"
	desc "github.com/Doremi203/couply/backend/payment/gen/api/payment-service/v1"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/payment-service"
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
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.GetPaymentStatusResponseToPB(response), nil
}
