package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSearch) CreateFilter(ctx context.Context, filter *search.Filter) error {
	query, args, err := buildCreateFilterQuery(filter)
	if err != nil {
		return errors.Wrapf(err, "buildCreateFilterQuery with %v", errors.Token("user_id", filter.UserID))
	}

	if err = executeCreateFilterQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateFilterQuery with %v", errors.Token("user_id", filter.UserID))
	}

	return nil
}

func buildCreateFilterQuery(filter *search.Filter) (string, []any, error) {
	query, args, err := sq.Insert(filtersTableName).
		Columns(filtersColumns...).
		Values(filter.UserID, filter.GenderPriority, filter.MinAge, filter.MaxAge, filter.MinHeight, filter.MaxHeight,
			filter.MinDistanceKM, filter.MaxDistanceKM, filter.Goal, filter.Zodiac, filter.Education, filter.Children,
			filter.Alcohol, filter.Smoking, filter.OnlyVerified, filter.OnlyPremium, filter.CreatedAt, filter.UpdatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateFilterQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsForeignKeyViolationError(err) {
			return user.ErrUserDoesntExist
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
