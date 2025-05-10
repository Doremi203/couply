package search

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (s *PgStorageSearch) GetFilter(ctx context.Context, userID uuid.UUID) (*search.Filter, error) {
	query, args, err := sq.Select(
		"user_id", "gender_priority", "min_age", "max_age", "min_height", "max_height",
		"min_distance_km", "max_distance_km", "goal", "zodiac", "education", "children", "alcohol", "smoking",
		"only_verified", "only_premium", "created_at", "updated_at",
	).
		From("filters").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	filter := &search.Filter{}
	err = s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...).Scan(
		&filter.UserID,
		&filter.GenderPriority,
		&filter.MinAge,
		&filter.MaxAge,
		&filter.MinHeight,
		&filter.MaxHeight,
		&filter.MinDistanceKM,
		&filter.MaxDistanceKM,
		&filter.Goal,
		&filter.Zodiac,
		&filter.Education,
		&filter.Children,
		&filter.Alcohol,
		&filter.Smoking,
		&filter.OnlyVerified,
		&filter.OnlyPremium,
		&filter.CreatedAt,
		&filter.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrFilterNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return filter, nil
}
