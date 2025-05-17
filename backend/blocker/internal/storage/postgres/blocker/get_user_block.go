package blocker

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrUserBlockNotFound = errors.Error("user block not found")
)

func (s *PgStorageBlocker) GetUserBlock(ctx context.Context, userID uuid.UUID) (*blocker.UserBlock, error) {
	query, args, err := sq.Select(
		"id", "blocked_id", "message", "created_at",
	).
		From("user_blocks").
		Where(sq.Eq{"blocked_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	b := &blocker.UserBlock{}
	err = s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...).Scan(
		&b.ID,
		&b.BlockedID,
		&b.Message,
		&b.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserBlockNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return b, nil
}
