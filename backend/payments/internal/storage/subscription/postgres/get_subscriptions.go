package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
)

type GetSubscriptionsOptions struct {
	SubscriptionStatus subscription.SubscriptionStatus
}

func (s *PgStorageSubscription) GetSubscriptions(ctx context.Context, opts GetSubscriptionsOptions) ([]*subscription.Subscription, error) {
	query, args, err := buildGetSubscriptionsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetSubscriptionsQuery with %v", errors.Token("options", opts))
	}

	subs, err := executeGetSubscriptionsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetSubscriptionsQuery with %v", errors.Token("options", opts))
	}

	return subs, nil
}

func buildGetSubscriptionsQuery(opts GetSubscriptionsOptions) (string, []any, error) {
	query, args, err := sq.Select(subscriptionsColumns...).
		From(subscriptionsTableName).
		Where(sq.Eq{
			statusColumn: opts.SubscriptionStatus,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetSubscriptionsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]*subscription.Subscription, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	subs, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[subscription.Subscription])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return subs, nil
}
