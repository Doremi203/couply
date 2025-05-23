package blocker

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageBlocker) AddUserBlock(ctx context.Context, block *blocker.UserBlock) error {
	query, args, err := sq.Insert("user_blocks").
		Columns(
			"id", "blocked_id", "message", "created_at", "status",
		).
		Values(
			block.ID, block.BlockedID, block.Message, block.CreatedAt, block.Status,
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
