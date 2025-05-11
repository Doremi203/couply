package push

import (
	"github.com/google/uuid"
)

type RecipientID uuid.UUID

func (id RecipientID) String() string {
	return uuid.UUID(id).String()
}

type Recipient struct {
	ID            RecipientID
	Subscriptions []Subscription
}
