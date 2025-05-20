package subscription_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/subscription-service"
	"github.com/Doremi203/couply/backend/payment/utils"
	"github.com/google/uuid"
)

func (c *UseCase) CreateSubscription(ctx context.Context, in *dto.CreateSubscriptionV1Request) (*dto.CreateSubscriptionV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	subscriptionID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newSubscription := &subscription.Subscription{
		ID:        subscriptionID,
		UserID:    userID,
		Plan:      in.GetSubscriptionPlan(),
		Status:    subscription.SubscriptionStatusPendingPayment,
		AutoRenew: in.GetAutoRenew(),
		StartDate: now,
		EndDate:   c.calculateEndDate(now, in.GetSubscriptionPlan()),
	}

	createdSubscription, err := c.subscriptionStorageFacade.CreateSubscriptionTx(ctx, newSubscription)
	if err != nil {
		return nil, err
	}

	return &dto.CreateSubscriptionV1Response{
		SubscriptionID:     createdSubscription.GetID(),
		SubscriptionPlan:   createdSubscription.GetPlan(),
		SubscriptionStatus: createdSubscription.GetStatus(),
		AutoRenew:          createdSubscription.GetAutoRenew(),
		StartDate:          createdSubscription.GetStartDate(),
		EndDate:            createdSubscription.GetEndDate(),
		PaymentIDs:         createdSubscription.GetPaymentIDs(),
	}, nil
}

func (c *UseCase) calculateEndDate(now time.Time, plan subscription.SubscriptionPlan) time.Time {
	durationMap := map[subscription.SubscriptionPlan]time.Duration{
		subscription.SubscriptionPlanMonthly:    30 * 24 * time.Hour,
		subscription.SubscriptionPlanSemiAnnual: 180 * 24 * time.Hour,
		subscription.SubscriptionPlanAnnual:     365 * 24 * time.Hour,
	}

	return now.Add(durationMap[plan])
}
