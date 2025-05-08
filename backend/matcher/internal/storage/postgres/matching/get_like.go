package matching

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrLikeNotFound = errors.Error("like not found")
)

func (s *PgStorageMatching) GetLike(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error) {
	query, args, err := sq.Select("sender_id", "receiver_id", "message", "created_at").
		From("likes").
		Where(sq.Eq{
			"sender_id":   senderID,
			"receiver_id": receiverID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetLike: failed to build query: %w", err)
	}

	row := s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...)

	like := &matching.Like{}
	err = row.Scan(
		&like.SenderID,
		&like.ReceiverID,
		&like.Message,
		&like.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrLikeNotFound
		}
		return nil, fmt.Errorf("GetLike: failed to scan row: %w", err)
	}

	return like, nil
}
