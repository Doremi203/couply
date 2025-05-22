package subscription

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *PgStorageSubscription) UpdateSubscriptionStatus(ctx context.Context, subID uuid.UUID, status subscription.SubscriptionStatus) error {
	query, args, err := buildUpdateSubscriptionStatusQuery(subID, status)
	if err != nil {
		return errors.Wrapf(err, "UpdateSubscriptionStatus with %v", errors.Token("subscription_id", subID))
	}

	result, err := executeUpdateSubscriptionStatusQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "UpdateSubscriptionStatus with %v", errors.Token("subscription_id", subID))
	}

	return verifyUpdateResult(result)
}

func buildUpdateSubscriptionStatusQuery(subID uuid.UUID, status subscription.SubscriptionStatus) (string, []any, error) {
	query, args, err := sq.Update("subscriptions").
		Set("status", status).
		Where(sq.Eq{
			"id": subID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildUpdateSubscriptionStatusQuery")
	}

	return query, args, nil
}

func executeUpdateSubscriptionStatusQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "executeUpdateSubscriptionStatusQuery")
	}
	return result, nil
}
