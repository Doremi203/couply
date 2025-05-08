package matching

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

var (
	ErrMatchAlreadyExists = errors.Error("match already exists between these users")
	ErrUserNotFound       = errors.Error("one or both users not found")
)

func (s *PgStorageMatching) AddMatch(ctx context.Context, match *matching.Match) error {
	user1, user2 := orderUserIDs(match.GetFirstUserID(), match.GetSecondUserID())

	query, args, err := sq.Insert("matches").
		Columns("first_user_id", "second_user_id", "created_at").
		Values(user1, user2, match.GetCreatedAt()).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("AddMatch: failed to build query: %w", err)
	}

	var (
		matchID   int64
		createdAt time.Time
	)

	err = s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...).Scan(
		&matchID,
		&createdAt,
	)
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
		return fmt.Errorf("AddMatch: %w", err)
	}

	return nil
}
