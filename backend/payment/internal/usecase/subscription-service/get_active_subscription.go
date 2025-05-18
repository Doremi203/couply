package subscription_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/subscription-service"
	"github.com/Doremi203/couply/backend/payment/utils"
)

func (c *UseCase) GetActiveSubscription(ctx context.Context, _ *dto.GetActiveSubscriptionV1Request) (*dto.GetActiveSubscriptionV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	sub, err := c.subscriptionStorageFacade.GetActiveSubscriptionTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.GetActiveSubscriptionV1Response{
		SubscriptionID:     sub.GetID(),
		SubscriptionPlan:   sub.GetPlan(),
		SubscriptionStatus: sub.GetStatus(),
		AutoRenew:          sub.GetAutoRenew(),
		StartDate:          sub.GetStartDate(),
		EndDate:            sub.GetEndDate(),
		PaymentIDs:         sub.GetPaymentIDs(),
	}, nil
}
