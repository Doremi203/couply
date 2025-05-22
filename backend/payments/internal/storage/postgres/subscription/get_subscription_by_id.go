package subscription

import (
	"context"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageSubscription) GetSubscriptionByID(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error) {
	query, args, err := buildSubscriptionByIDQuery(subID)
	if err != nil {
		return nil, errors.Wrapf(err, "GetSubscriptionByID with %v", errors.Token("subscription_id", subID))
	}

	sub, err := executeSubscriptionByIDQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "GetSubscriptionByID with %v", errors.Token("subscription_id", subID))
	}

	return sub, nil
}

func buildSubscriptionByIDQuery(subID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Select("id", "user_id", "plan", "status", "auto_renew", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{
			"id": subID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildSubscriptionByIDQuery")
	}
	return query, args, nil
}

func executeSubscriptionByIDQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (*subscription.Subscription, error) {
	row := queryEngine.QueryRow(ctx, query, args...)

	sub := &subscription.Subscription{}
	err := row.Scan(
		&sub.ID,
		&sub.UserID,
		&sub.Plan,
		&sub.Status,
		&sub.AutoRenew,
		&sub.StartDate,
		&sub.EndDate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errSubscriptionNotFound, "executeSubscriptionByIDQuery")
		}
		return nil, errors.Wrap(err, "executeSubscriptionByIDQuery")
	}
	return sub, nil
}
