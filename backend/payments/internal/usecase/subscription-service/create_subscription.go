package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/subscription-service"
)

func (c *UseCase) CreateSubscription(ctx context.Context, in *dto.CreateSubscriptionV1Request) (*dto.CreateSubscriptionV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSubscription")
	}

	newSubscription, err := dto.CreateSubscriptionRequestToSubscription(in, userID)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSubscription")
	}

	err = c.subscriptionStorageFacade.CreateSubscriptionTx(ctx, newSubscription)
	if err != nil {
		return nil, errors.Wrap(err, "CreateSubscription")
	}

	return dto.SubscriptionToCreateSubscriptionResponse(newSubscription), nil
}
