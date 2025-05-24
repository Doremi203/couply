package matching

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *PgStorageMatching) AddLike(ctx context.Context, like *matching.Like) error {
	query, args, err := sq.Insert("likes").
		Columns("sender_id", "receiver_id", "message", "status", "created_at").
		Values(
			like.SenderID,
			like.ReceiverID,
			like.Message,
			like.Status,
			like.CreatedAt,
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
			return ErrDuplicateLike
		}
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
