package blocker

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) UpdateUserBlockStatus(ctx context.Context, blockID uuid.UUID, status blocker.BlockStatus) error {
	query, args, err := sq.Update("user_blocks").
		Set("status", status).
		Where(sq.Eq{"id": blockID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user block status: %w", err)
	}

	return nil
}
