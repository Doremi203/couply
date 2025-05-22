package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) GetSubscriptionTx(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error) {
	sub, err := f.subscriptionStorage.GetSubscriptionByID(ctx, subID)
	if err != nil {
		return nil, errors.Wrap(err, "GetSubscriptionTx")
	}

	return sub, nil
}
