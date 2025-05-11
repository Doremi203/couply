package search

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageSearch) AddUserView(ctx context.Context, viewerID, viewedID uuid.UUID) error {
	query, args, err := sq.Insert("user_views").
		Columns("viewer_id", "viewed_id").
		Values(viewerID, viewedID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	query += " ON CONFLICT (viewer_id, viewed_id) DO NOTHING"

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert user view: %w", err)
	}

	return nil
}
