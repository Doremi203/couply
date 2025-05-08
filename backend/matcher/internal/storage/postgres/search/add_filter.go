package search

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *PgStorageSearch) AddFilter(ctx context.Context, filter *search.Filter) error {
	query, args, err := sq.Insert("filters").
		Columns(
			"user_id", "gender_priority", "min_age", "max_age", "min_height", "max_height",
			"distance", "goal", "zodiac", "education", "children", "alcohol", "smoking",
			"only_verified", "only_premium", "created_at", "updated_at",
		).
		Values(
			filter.GetUserID(),
			filter.GetGenderPriority(),
			filter.GetMinAge(),
			filter.GetMaxAge(),
			filter.GetMinHeight(),
			filter.GetMaxHeight(),
			filter.GetDistance(),
			filter.GetGoal(),
			filter.GetZodiac(),
			filter.GetEducation(),
			filter.GetChildren(),
			filter.GetAlcohol(),
			filter.GetSmoking(),
			filter.GetOnlyVerified(),
			filter.GetOnlyPremium(),
			filter.GetCreatedAt(),
			filter.GetUpdatedAt(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicateFilter
		}
		return fmt.Errorf("failed to insert filter: %w", err)
	}

	return nil
}
