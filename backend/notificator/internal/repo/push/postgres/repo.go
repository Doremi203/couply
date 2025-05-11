package pushpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/common/libs/slices"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/user"
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
		INSERT INTO push_subscriptions (endpoint, p256dh, auth_key, user_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (endpoint) DO UPDATE
		SET auth_key = EXCLUDED.auth_key,
		    user_id = EXCLUDED.user_id,
		    p256dh = EXCLUDED.p256dh;
	`
	_, err := r.db.Exec(
		ctx, query,
		subscription.Endpoint,
		subscription.Credentials.P256dh,
		subscription.Credentials.AuthKey,
		subscription.UserID,
	)
	if err != nil {
		return errors.WrapFail(err, "exec upsert push subscription query")
	}

	return nil
}

func (r *repo) DeleteSubscription(ctx context.Context, subscription push.Subscription) error {
	const query = `
		DELETE FROM push_subscriptions
		WHERE user_id = $1
	`
	_, err := r.db.Exec(
		ctx, query,
		subscription.UserID,
	)
	if err != nil {
		return errors.WrapFail(err, "exec delete push subscription query")
	}

	return nil
}

func (r *repo) GetSubscriptionsByUserID(ctx context.Context, userID user.ID) ([]push.Subscription, error) {
	const query = `
		SELECT
		    user_id,
		    endpoint,
		    p256dh,
		    auth_key
		FROM push_subscriptions
		WHERE user_id = $1
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
			UserID: user.ID(from.UserID),
		}
	}), nil
}
