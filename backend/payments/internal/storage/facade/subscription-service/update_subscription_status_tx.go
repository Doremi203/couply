package subscription_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionStatusTx(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error {
	err := f.subscriptionStorage.UpdateSubscriptionStatus(ctx, subscriptionID, status)
	if err != nil {
		return errors.Wrap(err, "UpdateSubscriptionStatusTx")
	}

	return nil
}
