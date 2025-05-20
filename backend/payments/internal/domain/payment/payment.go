package payment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID             uuid.UUID     `db:"id"`
	UserID         uuid.UUID     `db:"user_id"`
	SubscriptionID uuid.UUID     `db:"subscription_id"`
	Amount         int64         `db:"amount"`
	Currency       string        `db:"currency"`
	Status         PaymentStatus `db:"status"`
	GatewayID      string        `db:"gateway_id"`
	CreatedAt      time.Time     `db:"created_at"`
	UpdatedAt      time.Time     `db:"updated_at"`
}

func (x *Payment) GetID() uuid.UUID {
	if x != nil {
		return x.ID
	}
	return uuid.Nil
}

func (x *Payment) GetUserID() uuid.UUID {
	if x != nil {
		return x.UserID
	}
	return uuid.Nil
}

func (x *Payment) GetSubscriptionID() uuid.UUID {
	if x != nil {
		return x.SubscriptionID
	}
	return uuid.Nil
}

func (x *Payment) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Payment) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus(0)
}

func (x *Payment) GetGatewayID() string {
	if x != nil {
		return x.GatewayID
	}
	return ""
}

func (x *Payment) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func (x *Payment) GetUpdatedAt() time.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return time.Time{}
}
