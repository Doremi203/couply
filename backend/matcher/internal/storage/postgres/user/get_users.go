package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) GetUsers(ctx context.Context, userIDs []uuid.UUID) ([]*user.User, error) {
	if len(userIDs) == 0 {
		return []*user.User{}, nil
	}

	query, args, err := sq.Select(
		"id", "name", "age", "gender", "location", "bio", "goal", "zodiac",
		"height", "education", "children", "alcohol", "smoking", "is_hidden",
		"is_verified", "is_premium", "is_blocked", "created_at", "updated_at",
	).
		From("users").
		Where(sq.Eq{"id": userIDs}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build users query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute users query: %w", err)
	}
	defer rows.Close()

	users := make([]*user.User, 0)
	for rows.Next() {
		u := &user.User{}
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Age,
			&u.Gender,
			&u.Latitude,
			&u.Longitude,
			&u.BIO,
			&u.Goal,
			&u.Zodiac,
			&u.Height,
			&u.Education,
			&u.Children,
			&u.Alcohol,
			&u.Smoking,
			&u.IsHidden,
			&u.IsVerified,
			&u.IsPremium,
			&u.IsBlocked,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("users rows iteration error: %w", err)
	}

	return users, nil
}
