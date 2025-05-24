package postgres

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageMatching) UpdateLike(ctx context.Context, like *matching.Like) error {
	query, args, err := sq.Update("likes").
		Set("status", like.Status).
		Where(sq.Eq{
			"sender_id":   like.SenderID,
			"receiver_id": like.ReceiverID,
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
		return ErrLikeNotFound
	}

	return nil
}
