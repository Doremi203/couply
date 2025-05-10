package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) GetUser(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	query, args, err := sq.Select(
		"id", "name", "age", "gender", "latitude", "longitude", "bio", "goal", "zodiac",
		"height", "education", "children", "alcohol", "smoking", "is_hidden",
		"is_verified", "is_premium", "is_blocked", "created_at", "updated_at",
	).
		From("users").
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	u := &user.User{}
	err = s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return u, nil
}
