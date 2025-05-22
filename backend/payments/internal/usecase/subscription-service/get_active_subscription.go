package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
)

func (c *UseCase) GetActiveSubscription(ctx context.Context, _ *dto.GetActiveSubscriptionV1Request) (*dto.GetActiveSubscriptionV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetActiveSubscription")
	}

	sub, err := c.subscriptionStorageFacade.GetActiveSubscriptionTx(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetActiveSubscription")
	}

	return dto.SubscriptionToGetActiveSubscriptionResponse(sub), nil
}
