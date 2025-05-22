package subscription

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
)

var (
	errSubscriptionsNotFound = errors.Error("subscriptions not found")
)

func (s *PgStorageSubscription) GetSubscriptionsByStatus(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error) {
	query, args, err := buildSubscriptionsByStatusQuery(status)
	if err != nil {
		return nil, errors.Wrapf(err, "GetSubscriptionsByStatus with %v", status)
	}

	rows, err := executeSubscriptionsByStatusQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "GetSubscriptionsByStatus with %v", status)
	}
	defer rows.Close()

	return scanSubscriptions(rows)
}

func buildSubscriptionsByStatusQuery(status subscription.SubscriptionStatus) (string, []any, error) {
	query, args, err := sq.Select("id", "user_id", "plan", "status", "auto_renew", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{
			"status": status,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildSubscriptionsByStatusQuery")
	}
	return query, args, nil
}

func executeSubscriptionsByStatusQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgx.Rows, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errSubscriptionsNotFound, "executeSubscriptionsByStatusQuery")
		}
		return nil, errors.Wrap(err, "executeSubscriptionsByStatusQuery")
	}
	return rows, nil
}

func scanSubscriptions(rows pgx.Rows) ([]*subscription.Subscription, error) {
	var subscriptions []*subscription.Subscription

	for rows.Next() {
		sub := &subscription.Subscription{}
		err := rows.Scan(
			&sub.ID,
			&sub.UserID,
			&sub.Plan,
			&sub.Status,
			&sub.AutoRenew,
			&sub.StartDate,
			&sub.EndDate,
		)
		if err != nil {
			return nil, errors.Wrap(err, "scanSubscriptions")
		}
		subscriptions = append(subscriptions, sub)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scanSubscriptions")
	}

	return subscriptions, nil
}
