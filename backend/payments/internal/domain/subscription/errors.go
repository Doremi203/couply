package subscription

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrDuplicateSubscription = errors.Error("subscription already exists")
	ErrSubscriptionNotFound  = errors.Error("subscription not found")
	ErrSubscriptionsNotFound = errors.Error("subscriptions not found")

	ErrAlreadyActiveSubscriptionExists = errors.Error("already active subscription exists")
	ErrSubscriptionIsNotActive         = errors.Error("subscription is not active")
	ErrActiveSubscriptionDoesntExist   = errors.Error("active subscription does not exist")
	ErrSubscriptionHasAlreadyBeenPaid  = errors.Error("subscription has already been paid")
)
