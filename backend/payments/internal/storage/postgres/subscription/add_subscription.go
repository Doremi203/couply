package subscription

import (
	"context"

	pgerrors "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
)

var (
	errDuplicateSubscription = errors.Error("subscription already exists")
)

func (s *PgStorageSubscription) AddSubscription(ctx context.Context, subscription *subscription.Subscription) error {
	query, args, err := buildInsertSubscriptionQuery(subscription)
	if err != nil {
		return errors.Wrapf(err, "AddSubscription with %v", errors.Token("subscription", subscription))
	}

	if err = executeSubscriptionInsert(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "AddSubscription with %v", errors.Token("subscription", subscription))
	}

	return nil
}

func buildInsertSubscriptionQuery(sub *subscription.Subscription) (string, []any, error) {
	query, args, err := sq.Insert("subscriptions").
		Columns(
			"id", "user_id", "plan", "status",
			"start_date", "end_date", "auto_renew",
		).
		Values(
			sub.ID, sub.UserID, sub.Plan, sub.Status,
			sub.StartDate, sub.EndDate, sub.AutoRenew,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.WrapFail(err, "buildInsertSubscriptionQuery")
	}
	return query, args, nil
}

func executeSubscriptionInsert(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if pgerrors.IsUniqueViolationError(err) {
			return errors.Wrap(errDuplicateSubscription, "executeSubscriptionInsert")
		}
		return errors.Wrap(err, "executeSubscriptionInsert")
	}
	return nil
}
