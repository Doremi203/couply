package subscription_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/subscription-service"
)

func (c *UseCase) CancelSubscription(ctx context.Context, in *dto.CancelSubscriptionV1Request) (*dto.CancelSubscriptionV1Response, error) {
	err := c.subscriptionStorageFacade.CancelSubscriptionTx(ctx, in.GetSubscriptionID())
	if err != nil {
		return nil, err
	}

	return &dto.CancelSubscriptionV1Response{}, nil
}
