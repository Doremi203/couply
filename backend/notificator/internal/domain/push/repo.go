package push

import (
	"context"

	"github.com/Doremi203/couply/backend/notificator/internal/domain/user"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/push/repo_mock.go -typed

type Repo interface {
	UpsertSubscription(context.Context, Subscription) error
	DeleteSubscription(context.Context, Subscription) error
	GetSubscriptionsByUserID(context.Context, user.ID) ([]Subscription, error)
}
