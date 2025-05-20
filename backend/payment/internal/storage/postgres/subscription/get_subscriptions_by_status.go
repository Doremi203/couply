package subscription

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/payment/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSubscription) GetSubscriptionsByStatus(ctx context.Context, status subscription.SubscriptionStatus) ([]*subscription.Subscription, error) {
	query, args, err := sq.Select("id", "user_id", "plan", "status", "auto_renew", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{
			"status": status,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var subscriptions []*subscription.Subscription
	for rows.Next() {
		sub := &subscription.Subscription{}
		err = rows.Scan(
			&sub.ID,
			&sub.UserID,
			&sub.Plan,
			&sub.Status,
			&sub.AutoRenew,
			&sub.StartDate,
			&sub.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		subscriptions = append(subscriptions, sub)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(subscriptions) == 0 {
		return nil, errSubscriptionNotFound
	}

	return subscriptions, nil
}
