package search

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
)

func (s *PgStorageSearch) DeleteFilterInterests(ctx context.Context, id uuid.UUID) error {
	query, args, err := sq.Delete("filter_interests").
		Where(sq.Eq{"user_id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete filter interests: %w", err)
	}

	return nil
}
