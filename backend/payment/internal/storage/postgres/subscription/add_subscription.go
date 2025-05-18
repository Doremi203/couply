package subscription

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	errDuplicatedSubscription = errors.Error("subscription already exists for this user")
)

func (s *PgStorageSubscription) AddSubscription(ctx context.Context, subscription *subscription.Subscription) error {
	query, args, err := sq.Insert("subscriptions").
		Columns("id", "user_id", "plan", "status", "start_date", "end_date", "auto_renew").
		Values(
			subscription.GetID(),
			subscription.GetUserID(),
			subscription.GetPlan(),
			subscription.GetStatus(),
			subscription.GetStartDate(),
			subscription.GetEndDate(),
			subscription.GetAutoRenew(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errDuplicatedSubscription
		}
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
