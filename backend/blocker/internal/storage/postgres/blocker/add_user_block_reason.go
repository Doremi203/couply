package blocker

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) AddUserBlockReason(ctx context.Context, blockID uuid.UUID, reason blocker.ReportReason) error {
	query, args, err := sq.Insert("user_block_reasons").
		Columns(
			"block_id", "reason",
		).
		Values(
			blockID, reason,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert user block reason: %w", err)
	}

	return nil
}
