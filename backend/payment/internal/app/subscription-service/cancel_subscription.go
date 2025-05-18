package subscription_service

import (
	"context"
	desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/subscription-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CancelSubscription(ctx context.Context, in *desc.CancelSubscriptionV1Request) (*desc.CancelSubscriptionV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToCancelSubscriptionRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.CancelSubscription(ctx, req)
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.CancelSubscriptionResponseToPB(response), nil
}
