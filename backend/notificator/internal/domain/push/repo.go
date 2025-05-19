package push

import (
	"context"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/push/repo_mock.go -typed

type Repo interface {
	UpsertSubscription(context.Context, Subscription) error
	DeleteSubscription(context.Context, Subscription) error
	GetSubscriptionsByRecipientID(context.Context, RecipientID) ([]Subscription, error)
	GetAllRecipients(context.Context) ([]Recipient, error)
}
