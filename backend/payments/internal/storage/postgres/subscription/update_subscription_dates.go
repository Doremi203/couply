package subscription

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageSubscription) UpdateSubscriptionDates(ctx context.Context, subscriptionID uuid.UUID, startDate, endDate time.Time) error {
	query, args, err := sq.Update("subscriptions").
		Set("start_date", startDate).
		Set("end_date", endDate).
		Where(sq.Eq{
			"id": subscriptionID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	result, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errSubscriptionNotFound
	}

	return nil
}
