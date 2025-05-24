package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSearch) UpdateFilter(ctx context.Context, filter *search.Filter) error {
	query, args, err := buildUpdateFilterQuery(filter)
	if err != nil {
		return errors.Wrapf(err, "buildUpdateFilterQuery with %v", errors.Token("user_id", filter.UserID))
	}

	result, err := executeUpdateFilterQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdateUserQuery with %v", errors.Token("user_id", filter.UserID))
	}

	if result.RowsAffected() == 0 {
		return ErrFilterNotFound
	}

	return nil
}

func buildUpdateFilterQuery(filter *search.Filter) (string, []any, error) {
	query, args, err := sq.Update(filtersTableName).
		Set(genderPriorityColumnName, filter.GenderPriority).
		Set(minAgeColumnName, filter.MinAge).
		Set(maxAgeColumnName, filter.MaxAge).
		Set(minHeightColumnName, filter.MinHeight).
		Set(maxHeightColumnName, filter.MaxHeight).
		Set(minDistanceKMColumnName, filter.MinDistanceKM).
		Set(maxDistanceKMColumnName, filter.MaxDistanceKM).
		Set(goalColumnName, filter.Goal).
		Set(zodiacColumnName, filter.Zodiac).
		Set(educationColumnName, filter.Education).
		Set(childrenColumnName, filter.Children).
		Set(alcoholColumnName, filter.Alcohol).
		Set(smokingColumnName, filter.Smoking).
		Set(onlyPremiumColumnName, filter.OnlyVerified).
		Set(onlyPremiumColumnName, filter.OnlyPremium).
		Set(updatedAtColumnName, filter.UpdatedAt).
		Where(sq.Eq{userIdColumnName: filter.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdateFilterQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
