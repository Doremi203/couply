package subscription_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetActiveSubscriptionV1(ctx context.Context, in *desc.GetActiveSubscriptionV1Request) (*desc.GetActiveSubscriptionV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetActiveSubscription(ctx, dto.PBToGetActiveSubscriptionRequest(in))
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.GetActiveSubscriptionResponseToPB(response), nil
}
