package pushpostgres

import (
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/google/uuid"
)

type subscriptionEntity struct {
	Endpoint    string    `db:"endpoint"`
	P256dh      string    `db:"p256dh"`
	AuthKey     string    `db:"auth_key"`
	RecipientID uuid.UUID `db:"recipient_id"`
}

func entityToDomain(s subscriptionEntity) push.Subscription {
	return push.Subscription{
		RecipientID: push.RecipientID(s.RecipientID),
		Endpoint:    push.Endpoint(s.Endpoint),
		Credentials: push.Credentials{
			P256dh:  s.P256dh,
			AuthKey: s.AuthKey,
		},
	}
}
