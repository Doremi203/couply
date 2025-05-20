package subscription_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateSubscriptionV1(ctx context.Context, in *desc.CreateSubscriptionV1Request) (*desc.CreateSubscriptionV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.CreateSubscription(ctx, dto.PBToCreateSubscriptionRequest(in))
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.CreateSubscriptionResponseToPB(response), nil
}
