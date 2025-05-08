package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query, args, err := sq.Delete("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
