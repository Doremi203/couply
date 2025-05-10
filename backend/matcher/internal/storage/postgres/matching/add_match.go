package matching

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (s *PgStorageMatching) AddMatch(ctx context.Context, match *matching.Match) error {
	user1, user2 := orderUserIDs(match.GetFirstUserID(), match.GetSecondUserID())

	query, args, err := sq.Insert("matches").
		Columns("first_user_id", "second_user_id", "created_at").
		Values(user1, user2, match.GetCreatedAt()).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return ErrMatchAlreadyExists
			case "23503":
				return ErrUserNotFound
			}
		}
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

// orderUserIDs гарантирует один порядок айди пользователей
func orderUserIDs(id1, id2 uuid.UUID) (uuid.UUID, uuid.UUID) {
	if bytes.Compare(id1[:], id2[:]) < 0 {
		return id1, id2
	}
	return id2, id1
}
