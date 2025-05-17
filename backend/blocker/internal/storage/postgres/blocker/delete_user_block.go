package blocker

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) DeleteUserBlock(ctx context.Context, userID uuid.UUID) error {
	query, args, err := sq.Delete("user_blocks").
		Where(sq.Eq{"blocked_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user block: %w", err)
	}

	return nil
}
