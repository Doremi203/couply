package subscription

import desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"

type SubscriptionPlan int

const (
	SubscriptionPlanUnspecified SubscriptionPlan = iota
	SubscriptionPlanMonthly
	SubscriptionPlanSemiAnnual
	SubscriptionPlanAnnual
)

func PBToSubscriptionPlan(subscriptionPlan desc.SubscriptionPlan) SubscriptionPlan {
	switch subscriptionPlan {
	case desc.SubscriptionPlan_SUBSCRIPTION_PLAN_UNSPECIFIED:
		return SubscriptionPlanUnspecified
	case desc.SubscriptionPlan_SUBSCRIPTION_PLAN_MONTHLY:
		return SubscriptionPlanMonthly
	case desc.SubscriptionPlan_SUBSCRIPTION_PLAN_SEMI_ANNUAL:
		return SubscriptionPlanSemiAnnual
	case desc.SubscriptionPlan_SUBSCRIPTION_PLAN_ANNUAL:
		return SubscriptionPlanAnnual
	default:
		return SubscriptionPlan(0)
	}
}

func SubscriptionPlanToPB(subscriptionPlan SubscriptionPlan) desc.SubscriptionPlan {
	switch subscriptionPlan {
	case SubscriptionPlanUnspecified:
		return desc.SubscriptionPlan_SUBSCRIPTION_PLAN_UNSPECIFIED
	case SubscriptionPlanMonthly:
		return desc.SubscriptionPlan_SUBSCRIPTION_PLAN_MONTHLY
	case SubscriptionPlanSemiAnnual:
		return desc.SubscriptionPlan_SUBSCRIPTION_PLAN_SEMI_ANNUAL
	case SubscriptionPlanAnnual:
		return desc.SubscriptionPlan_SUBSCRIPTION_PLAN_ANNUAL
	default:
		return desc.SubscriptionPlan(0)
	}
}
