package matching

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageMatching) FetchIncomingLikes(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error) {
	query, args, err := sq.Select(
		"sender_id",
		"receiver_id",
		"message",
		"status",
		"created_at",
	).
		From("likes").
		Where(sq.Eq{"receiver_id": userID}).
		OrderBy("created_at DESC").
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("FetchIncomingLikes: failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("FetchIncomingLikes: query error: %w", err)
	}
	defer rows.Close()

	var likes []*matching.Like
	for rows.Next() {
		var like matching.Like
		err := rows.Scan(
			&like.SenderID,
			&like.ReceiverID,
			&like.Message,
			&like.Status,
			&like.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("FetchIncomingLikes: scan error: %w", err)
		}
		likes = append(likes, &like)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchIncomingLikes: rows error: %w", err)
	}

	return likes, nil
}
