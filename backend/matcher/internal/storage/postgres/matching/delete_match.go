package matching

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageMatching) DeleteMatch(ctx context.Context, userID, targetUserID uuid.UUID) error {
	user1, user2 := orderUserIDs(userID, targetUserID)

	query, args, err := sq.Delete("matches").
		Where(sq.Eq{"first_user_id": user1}).
		Where(sq.Eq{"second_user_id": user2}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete match: %w", err)
	}

	return nil
}
