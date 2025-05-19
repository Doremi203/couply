package pushpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/common/libs/slices"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func NewRepo(db postgres.Client) *repo {
	return &repo{db: db}
}

type repo struct {
	db postgres.Client
}

func (r *repo) UpsertSubscription(ctx context.Context, subscription push.Subscription) error {
	const query = `
		INSERT INTO push_subscriptions (endpoint, p256dh, auth_key, recipient_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (endpoint) DO UPDATE
		SET auth_key = EXCLUDED.auth_key,
		    recipient_id = EXCLUDED.recipient_id,
		    p256dh = EXCLUDED.p256dh;
	`
	_, err := r.db.Exec(
		ctx, query,
		subscription.Endpoint,
		subscription.Credentials.P256dh,
		subscription.Credentials.AuthKey,
		subscription.RecipientID,
	)
	if err != nil {
		return errors.WrapFail(err, "exec upsert push subscription query")
	}

	return nil
}

func (r *repo) DeleteSubscription(ctx context.Context, subscription push.Subscription) error {
	const query = `
		DELETE FROM push_subscriptions
		WHERE endpoint = $1 AND recipient_id = $2
	`
	_, err := r.db.Exec(
		ctx, query,
		subscription.Endpoint,
		subscription.RecipientID,
	)
	if err != nil {
		return errors.WrapFail(err, "exec delete push subscription query")
	}

	return nil
}

func (r *repo) GetSubscriptionsByRecipientID(ctx context.Context, userID push.RecipientID) ([]push.Subscription, error) {
	const query = `
		SELECT
		    recipient_id,
		    endpoint,
		    p256dh,
		    auth_key
		FROM push_subscriptions
		WHERE recipient_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, errors.WrapFail(err, "exec select push subscriptions query")
	}
	defer rows.Close()

	subscriptionEntities, err := pgx.CollectRows(rows, pgx.RowToStructByName[subscriptionEntity])
	if err != nil {
		return nil, errors.WrapFail(err, "collect push subscriptions rows")
	}

	return slices.Map(subscriptionEntities, func(from subscriptionEntity) push.Subscription {
		return push.Subscription{
			Endpoint: push.Endpoint(from.Endpoint),
			Credentials: push.Credentials{
				P256dh:  from.P256dh,
				AuthKey: from.AuthKey,
			},
			RecipientID: push.RecipientID(from.RecipientID),
		}
	}), nil
}

func (r *repo) GetAllRecipients(ctx context.Context) ([]push.Recipient, error) {
	const query = `
		SELECT 
		    endpoint,
		    p256dh,
		    auth_key,
		    recipient_id
		FROM push_subscriptions
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, errors.WrapFail(err, "exec select all push subscriptions query")
	}
	defer rows.Close()

	subscriptions, err := pgx.CollectRows(rows, pgx.RowToStructByName[subscriptionEntity])
	if err != nil {
		return nil, errors.WrapFail(err, "collect all push subscriptions rows")
	}

	subscriptionsMap := make(map[uuid.UUID][]subscriptionEntity, len(subscriptions))

	for i := range subscriptions {
		subscriptionsMap[subscriptions[i].RecipientID] = append(subscriptionsMap[subscriptions[i].RecipientID], subscriptions[i])
	}

	recipients := make([]push.Recipient, 0, len(subscriptionsMap))
	for recipientID, subscriptions := range subscriptionsMap {
		recipients = append(recipients, push.Recipient{
			ID:            push.RecipientID(recipientID),
			Subscriptions: slices.Map(subscriptions, entityToDomain),
		})
	}

	return recipients, nil
}
