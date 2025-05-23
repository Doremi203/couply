package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
)

func (c *UseCase) CancelSubscription(ctx context.Context, in *dto.CancelSubscriptionV1Request) (*dto.CancelSubscriptionV1Response, error) {
	err := c.subscriptionStorageFacade.CancelSubscriptionTx(ctx, in.SubscriptionID)
	if err != nil {
		return nil, errors.Wrap(err, "CancelSubscription")
	}

	return &dto.CancelSubscriptionV1Response{}, nil
}
