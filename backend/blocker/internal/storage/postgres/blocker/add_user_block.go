package blocker

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) AddUserBlock(ctx context.Context, blockID, blockedUserID uuid.UUID, message string, createdAt time.Time) error {
	query, args, err := sq.Insert("user_blocks").
		Columns(
			"id", "blocked_id", "message", "created_at",
		).
		Values(
			blockID, blockedUserID, message, createdAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert user block: %w", err)
	}

	return nil
}
