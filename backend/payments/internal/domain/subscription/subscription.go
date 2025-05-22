package subscription

import (
	"time"

	"github.com/google/uuid"
)

const (
	monthDays    = 30
	halfYearDays = 180
	yearDays     = 365

	monthPrice      = 199
	semiAnnualPrice = 999
	annualPrice     = 1799
)

type Subscription struct {
	ID         uuid.UUID          `db:"id"`
	UserID     uuid.UUID          `db:"user_id"`
	Plan       SubscriptionPlan   `db:"plan"`
	Status     SubscriptionStatus `db:"status"`
	AutoRenew  bool               `db:"auto_renew"`
	StartDate  time.Time          `db:"start_date"`
	EndDate    time.Time          `db:"end_date"`
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

func CalculateEndDate(now time.Time, plan SubscriptionPlan) time.Time {
	durationMap := map[SubscriptionPlan]time.Duration{
		SubscriptionPlanMonthly:    monthDays * 24 * time.Hour,
		SubscriptionPlanSemiAnnual: halfYearDays * 24 * time.Hour,
		SubscriptionPlanAnnual:     yearDays * 24 * time.Hour,
	}
	return now.Add(durationMap[plan])
}

func GetPlanPrice(plan SubscriptionPlan) int64 {
	planPrices := map[SubscriptionPlan]int64{
		SubscriptionPlanMonthly:    monthPrice,
		SubscriptionPlanSemiAnnual: semiAnnualPrice,
		SubscriptionPlanAnnual:     annualPrice,
	}
	return planPrices[plan]
}
