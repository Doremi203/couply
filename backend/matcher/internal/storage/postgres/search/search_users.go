package search

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageSearch) SearchUsers(
	ctx context.Context,
	filter *search.Filter,
	interests *interest.Interest,
	offset, limit uint64,
) ([]*user.User, error) {
	query, args, err := buildSearchQuery(filter, interests, offset, limit)
	if err != nil {
		return nil, errors.WrapFail(err, "build query")
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, errors.WrapFail(err, "query")
	}
	defer rows.Close()

	return scanUsers(rows)
}

func scanUsers(rows pgx.Rows) ([]*user.User, error) {
	var users []*user.User
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, errors.WrapFail(err, "scan rows")
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func scanUser(row pgx.Row) (*user.User, error) {
	var u user.User
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Age,
		&u.Gender,
		&u.Location,
		&u.BIO,
		&u.Goal,
		&u.Zodiac,
		&u.Height,
		&u.Education,
		&u.Children,
		&u.Alcohol,
		&u.Smoking,
		&u.Hidden,
		&u.Verified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	return &u, err
}

func buildSearchQuery(
	filter *search.Filter,
	interests *interest.Interest,
	offset, limit uint64,
) (string, []any, error) {
	qb := baseQuery().Where(baseConditions(filter))

	qb = applyMainFilters(qb, filter)
	qb = applyInterestFilters(qb, interests)
	qb = applyPagination(qb, offset, limit)

	return qb.ToSql()
}

func baseQuery() sq.SelectBuilder {
	return sq.Select(
		"id", "name", "age", "gender", "location",
		"bio", "goal", "zodiac", "height", "education",
		"children", "alcohol", "smoking", "hidden",
		"verified", "created_at", "updated_at",
	).From("users u").
		PlaceholderFormat(sq.Dollar)
}

func baseConditions(filter *search.Filter) sq.Sqlizer {
	return sq.And{
		sq.Eq{"hidden": false},
		sq.NotEq{"id": filter.GetUserID()},
	}
}

func applyMainFilters(qb sq.SelectBuilder, filter *search.Filter) sq.SelectBuilder {
	qb = applyRangeFilter(qb, "age", filter.GetMinAge(), filter.GetMaxAge())
	qb = applyRangeFilter(qb, "height", filter.GetMinHeight(), filter.GetMaxHeight())

	if filter.GetGenderPriority() != 3 { // 3 - ANY
		qb = qb.Where(sq.Eq{"gender": int(filter.GetGenderPriority())})
	}

	filters := map[string]int{
		"goal":      int(filter.GetGoal()),
		"zodiac":    int(filter.GetZodiac()),
		"education": int(filter.GetEducation()),
		"children":  int(filter.GetChildren()),
		"alcohol":   int(filter.GetAlcohol()),
		"smoking":   int(filter.GetSmoking()),
	}

	for field, value := range filters {
		if value != 0 {
			qb = qb.Where(sq.Eq{field: value})
		}
	}

	return qb
}

func applyRangeFilter(
	qb sq.SelectBuilder,
	field string,
	min, max int32,
) sq.SelectBuilder {
	if min > 0 && max >= min {
		qb = qb.Where(sq.And{
			sq.GtOrEq{field: min},
			sq.LtOrEq{field: max},
		})
	}
	return qb
}

func applyInterestFilters(qb sq.SelectBuilder, interests *interest.Interest) sq.SelectBuilder {
	filterInterests := extractInterestPairs(interests)
	if len(filterInterests) == 0 {
		return qb
	}

	sub := sq.Select("i.user_id").
		From("interests i").
		Where(buildInterestConditions(filterInterests)).
		GroupBy("i.user_id").
		Having("COUNT(DISTINCT i.type, i.value) = ?", len(filterInterests))

	qb.Where(sq.Expr("EXISTS (?)", sub))

	return qb
}

func extractInterestPairs(interests *interest.Interest) []struct {
	Type  string
	Value int
} {
	interestGroups := map[string][]int{
		"social":           convertSlice(interests.GetSocial()),
		"sport":            convertSlice(interests.GetSport()),
		"self_development": convertSlice(interests.GetSelfDevelopment()),
		"art":              convertSlice(interests.GetArt()),
		"hobby":            convertSlice(interests.GetHobby()),
		"gastronomy":       convertSlice(interests.GetGastronomy()),
	}

	var pairs []struct {
		Type  string
		Value int
	}
	for t, values := range interestGroups {
		for _, v := range values {
			if v != 0 {
				pairs = append(pairs, struct {
					Type  string
					Value int
				}{t, v})
			}
		}
	}
	return pairs
}

func buildInterestConditions(pairs []struct {
	Type  string
	Value int
}) sq.Sqlizer {
	var conditions []sq.Sqlizer
	for _, p := range pairs {
		conditions = append(conditions, sq.And{
			sq.Eq{"type": p.Type},
			sq.Eq{"value": p.Value},
		})
	}
	return sq.Or(conditions)
}

func applyPagination(qb sq.SelectBuilder, offset, limit uint64) sq.SelectBuilder {
	if limit > 0 {
		qb = qb.Limit(limit)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	return qb.OrderBy("created_at DESC")
}
