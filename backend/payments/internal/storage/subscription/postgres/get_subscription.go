package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetSubscriptionOptions struct {
	SubscriptionID              uuid.UUID
	UserID                      uuid.UUID
	ActiveSubscription          bool
	ActiveOrPendingSubscription bool
}

func (s *PgStorageSubscription) GetSubscription(ctx context.Context, opts GetSubscriptionOptions) (*subscription.Subscription, error) {
	query, args, err := buildGetSubscriptionQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetSubscriptionQuery with %v", errors.Token("options", opts))
	}

	sub, err := executeGetSubscriptionQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetSubscriptionQuery with %v", errors.Token("options", opts))
	}

	return sub, nil
}

func buildGetSubscriptionQuery(opts GetSubscriptionOptions) (string, []any, error) {
	sb := sq.Select(subscriptionsColumns...).
		From(subscriptionsTableName)

	if opts.ActiveSubscription {
		sb = sb.Where(sq.Eq{
			userIDColumn: opts.UserID,
			statusColumn: subscription.SubscriptionStatusActive,
		})
	} else if opts.ActiveOrPendingSubscription {
		sb = sb.Where(sq.Eq{userIDColumn: opts.UserID}).
			Where(sq.Or{
				sq.Eq{statusColumn: subscription.SubscriptionStatusActive},
				sq.Eq{statusColumn: subscription.SubscriptionStatusPendingPayment},
			})
	} else {
		sb = sb.Where(sq.Eq{idColumn: opts.SubscriptionID})
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetSubscriptionQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*subscription.Subscription, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	sub, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[subscription.Subscription])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(subscription.ErrSubscriptionNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return sub, nil
}
