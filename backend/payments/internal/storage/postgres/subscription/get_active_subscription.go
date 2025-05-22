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

var (
	errSubscriptionNotFound = errors.Error("subscription not found")
)

func (s *PgStorageSubscription) GetActiveSubscriptionByUserID(ctx context.Context, userID uuid.UUID) (*subscription.Subscription, error) {
	query, args, err := buildActiveSubscriptionQuery(userID)
	if err != nil {
		return nil, errors.Wrapf(err, "GetActiveSubscriptionByUserID with %v", errors.Token("user_id", userID))
	}

	sub, err := executeActiveSubscriptionQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "GetActiveSubscriptionByUserID with %v", errors.Token("user_id", userID))
	}

	return sub, nil
}

func buildActiveSubscriptionQuery(userID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Select("id", "user_id", "plan", "status", "auto_renew", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{
			"user_id": userID,
			"status":  subscription.SubscriptionStatusActive,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildActiveSubscriptionQuery")
	}
	return query, args, nil
}

func executeActiveSubscriptionQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (*subscription.Subscription, error) {
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
			return nil, errors.Wrap(errSubscriptionNotFound, "executeActiveSubscriptionQuery")
		}
		return nil, errors.Wrap(err, "executeActiveSubscriptionQuery")
	}
	return sub, nil
}
