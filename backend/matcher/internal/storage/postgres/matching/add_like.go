package matching

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrDuplicateLike = errors.Error("like already exists between these users")
)

func (s *PgStorageMatching) AddLike(ctx context.Context, like *matching.Like) error {
	query, args, err := sq.Insert("likes").
		Columns("sender_id", "receiver_id", "message", "status", "created_at").
		Values(
			like.GetSenderID(),
			like.GetReceiverID(),
			like.GetMessage(),
			like.GetStatus(),
			like.GetCreatedAt(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("AddLike: failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicateLike
		}
		return fmt.Errorf("AddLike: %w", err)
	}

	return nil
}
