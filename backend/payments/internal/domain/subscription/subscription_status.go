package subscription

import desc "github.com/Doremi203/couply/backend/payments/gen/api/subscription-service/v1"

type SubscriptionStatus int

const (
	SubscriptionStatusUnspecified SubscriptionStatus = iota
	SubscriptionStatusActive
	SubscriptionStatusExpired
	SubscriptionStatusCanceled
	SubscriptionStatusPendingPayment
)

func PBToSubscriptionStatus(subscriptionStatus desc.SubscriptionStatus) SubscriptionStatus {
	switch subscriptionStatus {
	case desc.SubscriptionStatus_SUBSCRIPTION_STATUS_UNSPECIFIED:
		return SubscriptionStatusUnspecified
	case desc.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE:
		return SubscriptionStatusActive
	case desc.SubscriptionStatus_SUBSCRIPTION_STATUS_EXPIRED:
		return SubscriptionStatusExpired
	case desc.SubscriptionStatus_SUBSCRIPTION_STATUS_CANCELED:
		return SubscriptionStatusCanceled
	case desc.SubscriptionStatus_SUBSCRIPTION_STATUS_PENDING_PAYMENT:
		return SubscriptionStatusPendingPayment
	default:
		return SubscriptionStatus(0)
	}
}

func SubscriptionStatusToPB(subscriptionStatus SubscriptionStatus) desc.SubscriptionStatus {
	switch subscriptionStatus {
	case SubscriptionStatusUnspecified:
		return desc.SubscriptionStatus_SUBSCRIPTION_STATUS_UNSPECIFIED
	case SubscriptionStatusActive:
		return desc.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE
	case SubscriptionStatusExpired:
		return desc.SubscriptionStatus_SUBSCRIPTION_STATUS_EXPIRED
	case SubscriptionStatusCanceled:
		return desc.SubscriptionStatus_SUBSCRIPTION_STATUS_CANCELED
	case SubscriptionStatusPendingPayment:
		return desc.SubscriptionStatus_SUBSCRIPTION_STATUS_PENDING_PAYMENT
	default:
		return desc.SubscriptionStatus(0)
	}
}
