package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"

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
	switch {
	case errors.Is(err, subscription.ErrAlreadyActiveSubscriptionExists):
		return nil, status.Error(codes.AlreadyExists, subscription.ErrAlreadyActiveSubscriptionExists.Error())
	case errors.Is(err, subscription.ErrDuplicateSubscription):
		return nil, status.Error(codes.AlreadyExists, subscription.ErrDuplicateSubscription.Error())
	case err != nil:
		return nil, err
	}

	return dto.CreateSubscriptionResponseToPB(response), nil
}
