package subscription

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/payments/internal/domain/subscription"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageSubscription) UpdateSubscriptionStatus(ctx context.Context, subscriptionID uuid.UUID, status subscription.SubscriptionStatus) error {
	query, args, err := sq.Update("subscriptions").
		Set("status", status).
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
