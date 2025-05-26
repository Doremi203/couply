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

func (i *Implementation) CancelSubscriptionV1(ctx context.Context, in *desc.CancelSubscriptionV1Request) (*desc.CancelSubscriptionV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToCancelSubscriptionRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.CancelSubscription(ctx, req)
	switch {
	case errors.Is(err, subscription.ErrSubscriptionNotFound):
		return nil, status.Error(codes.NotFound, subscription.ErrSubscriptionNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.CancelSubscriptionResponseToPB(response), nil
}
