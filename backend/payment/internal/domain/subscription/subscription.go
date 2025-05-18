package subscription

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Plan       SubscriptionPlan
	Status     SubscriptionStatus
	AutoRenew  bool
	StartDate  time.Time
	EndDate    time.Time
	PaymentIDs []uuid.UUID
}

func (x *Subscription) GetID() uuid.UUID {
	if x != nil {
		return x.ID
	}
	return uuid.Nil
}

func (x *Subscription) GetUserID() uuid.UUID {
	if x != nil {
		return x.UserID
	}
	return uuid.Nil
}

func (x *Subscription) GetPlan() SubscriptionPlan {
	if x != nil {
		return x.Plan
	}
	return SubscriptionPlan(0)
}

func (x *Subscription) GetStatus() SubscriptionStatus {
	if x != nil {
		return x.Status
	}
	return SubscriptionStatus(0)
}

func (x *Subscription) GetAutoRenew() bool {
	if x != nil {
		return x.AutoRenew
	}
	return false
}

func (x *Subscription) GetStartDate() time.Time {
	if x != nil {
		return x.StartDate
	}
	return time.Time{}
}

func (x *Subscription) GetEndDate() time.Time {
	if x != nil {
		return x.EndDate
	}
	return time.Time{}
}

func (x *Subscription) GetPaymentIDs() []uuid.UUID {
	if x != nil {
		return x.PaymentIDs
	}
	return nil
}
