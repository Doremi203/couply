package push

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

type Endpoint string

type Credentials struct {
	P256dh  string
	AuthKey string
}

func NewSubscription(
	recipientID RecipientID,
	endpoint string,
	p256dh string,
	authKey string,
) (Subscription, error) {
	if endpoint == "" {
		return Subscription{}, errors.Error("push endpoint should not be empty")
	}
	if p256dh == "" {
		return Subscription{}, errors.Error("push public key should not be empty")
	}
	if authKey == "" {
		return Subscription{}, errors.Error("push secret key should not be empty")
	}

	return Subscription{
		Endpoint: Endpoint(endpoint),
		Credentials: Credentials{
			P256dh:  p256dh,
			AuthKey: authKey,
		},
		RecipientID: recipientID,
	}, nil
}

type Subscription struct {
	RecipientID RecipientID
	Endpoint    Endpoint
	Credentials Credentials
}
