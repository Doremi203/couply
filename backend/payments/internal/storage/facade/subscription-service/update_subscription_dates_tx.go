package subscription_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeSubscription) UpdateSubscriptionDatesTx(ctx context.Context, subscriptionID uuid.UUID, startDate, endDate time.Time) error {
	err := f.subscriptionStorage.UpdateSubscriptionDatesByID(ctx, subscriptionID, startDate, endDate)
	if err != nil {
		return errors.Wrap(err, "UpdateSubscriptionDatesTx")
	}

	return nil
}
