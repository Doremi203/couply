package subscription

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageSubscription) GetSubscription(ctx context.Context, subID uuid.UUID) (*subscription.Subscription, error) {
	query, args, err := sq.Select("id", "user_id", "plan", "status", "auto_renew", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{
			"id": subID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	row := s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...)

	sub := &subscription.Subscription{}
	err = row.Scan(
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
			return nil, errSubscriptionNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return sub, nil
}
