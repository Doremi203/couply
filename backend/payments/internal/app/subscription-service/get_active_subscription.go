package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"

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
	switch {
	case errors.Is(err, subscription.ErrActiveSubscriptionDoesntExist):
		return nil, status.Error(codes.FailedPrecondition, subscription.ErrActiveSubscriptionDoesntExist.Error())
	case errors.Is(err, payment.ErrPaymentsNotFound):
		return nil, status.Error(codes.NotFound, payment.ErrPaymentsNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.GetActiveSubscriptionResponseToPB(response), nil
}
