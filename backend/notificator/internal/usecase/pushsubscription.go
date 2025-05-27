package usecase

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
)

func NewPushSubscription(pushRepo push.Repo) PushSubscription {
	return PushSubscription{pushRepo: pushRepo}
}

type PushSubscription struct {
	pushRepo push.Repo
}

func (s PushSubscription) Subscribe(ctx context.Context, subscription push.Subscription) error {
	err := s.pushRepo.UpsertSubscription(ctx, subscription)
	if err != nil {
		return errors.WrapFailf(
			err,
			"upsert push subscription for recipient with %v",
			errors.Token("id", subscription.RecipientID),
		)
	}

	return nil
}

func (s PushSubscription) Unsubscribe(ctx context.Context, subscription push.Subscription) error {
	err := s.pushRepo.DeleteSubscription(ctx, subscription)
	if err != nil {
		return errors.WrapFailf(
			err,
			"delete push subscription for recipient with %v",
			errors.Token("id", subscription.RecipientID),
		)
	}

	return nil
}

func (s PushSubscription) GetRecipient(ctx context.Context, recipientID push.RecipientID) (push.Recipient, error) {
	subs, err := s.pushRepo.GetSubscriptionsByRecipientID(ctx, recipientID)
	if err != nil {
		return push.Recipient{}, errors.WrapFailf(
			err,
			"delete push subscription for recipient with %v",
			errors.Token("id", recipientID),
		)
	}

	return push.Recipient{
		ID:            recipientID,
		Subscriptions: subs,
	}, nil
}
