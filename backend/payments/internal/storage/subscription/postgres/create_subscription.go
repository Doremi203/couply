package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	pgerrors "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSubscription) CreateSubscription(ctx context.Context, subscription *subscription.Subscription) error {
	query, args, err := buildCreateSubscriptionQuery(subscription)
	if err != nil {
		return errors.Wrapf(err, "buildCreateSubscriptionQuery with %v", errors.Token("subscription", subscription))
	}

	if err = executeCreateSubscription(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateSubscription with %v", errors.Token("subscription", subscription))
	}

	return nil
}

func buildCreateSubscriptionQuery(sub *subscription.Subscription) (string, []any, error) {
	query, args, err := sq.Insert(subscriptionsTableName).
		Columns(subscriptionsColumns...).
		Values(sub.ID, sub.UserID, sub.Plan, sub.Status, sub.StartDate, sub.EndDate, sub.AutoRenew).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateSubscription(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if pgerrors.IsUniqueViolationError(err) {
			return errors.Wrap(subscription.ErrDuplicateSubscription, "exec")
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
