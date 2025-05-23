package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
)

func (f *StorageFacadeSubscription) CreateSubscriptionTx(ctx context.Context, newSubscription *subscription.Subscription) error {
	err := f.subscriptionStorage.CreateSubscription(ctx, newSubscription)
	if err != nil {
		return errors.Wrap(err, "storage.CreateSubscription")
	}

	return nil
}
