package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
)

func (c *UseCase) CancelSubscription(ctx context.Context, in *dto.CancelSubscriptionV1Request) (*dto.CancelSubscriptionV1Response, error) {
	err := c.subscriptionStorageFacade.UpdateSubscriptionStatusTx(ctx, in.GetSubscriptionID(), subscription.SubscriptionStatusCanceled)
	if err != nil {
		return nil, err
	}

	return &dto.CancelSubscriptionV1Response{}, nil
}
