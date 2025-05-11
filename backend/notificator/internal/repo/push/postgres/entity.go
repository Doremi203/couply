package pushpostgres

import (
	"github.com/google/uuid"
)

type subscriptionEntity struct {
	Endpoint    string    `db:"endpoint"`
	P256dh      string    `db:"p256dh"`
	AuthKey     string    `db:"auth_key"`
	RecipientID uuid.UUID `db:"recipient_id"`
}
