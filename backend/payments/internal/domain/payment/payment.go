package payment

import (
	"time"

	"github.com/google/uuid"
)

const (
	// MainCurrency for payments
	MainCurrency = "RUB"
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
