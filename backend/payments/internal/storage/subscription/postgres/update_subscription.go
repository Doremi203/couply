package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	"github.com/Doremi203/couply/backend/payments/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"

	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSubscription) UpdateSubscription(ctx context.Context, sub *subscription.Subscription) error {
	query, args, err := buildUpdateSubscriptionQuery(sub)
	if err != nil {
		return errors.Wrapf(err, "buildUpdateSubscriptionQuery with %v", errors.Token("subscription", sub))
	}

	result, err := executeUpdateSubscriptionQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdateSubscriptionQuery with %v", errors.Token("subscription", sub))
	}

	if result.RowsAffected() == 0 {
		return subscription.ErrSubscriptionNotFound
	}

	return nil
}

func buildUpdateSubscriptionQuery(sub *subscription.Subscription) (string, []any, error) {
	query, args, err := sq.Update(subscriptionsTableName).
		Set(statusColumn, sub.Status).
		Set(startDateColumn, sub.StartDate).
		Set(endDateColumn, sub.EndDate).
		Where(sq.Eq{
			idColumn: sub.ID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "update")
	}
	return query, args, nil
}

func executeUpdateSubscriptionQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
