package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

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
	curLatitude, curLongitude float64,
	offset, limit uint64,
) ([]*user.User, map[uuid.UUID]float64, error) {
	query, args, err := buildSearchQuery(filter, interests, curLatitude, curLongitude, offset, limit)
	if err != nil {
		return nil, nil, errors.WrapFail(err, "build query")
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, nil, errors.WrapFail(err, "query")
	}
	defer rows.Close()

	return scanUsers(rows)
}

func scanUsers(rows pgx.Rows) ([]*user.User, map[uuid.UUID]float64, error) {
	var users []*user.User
	distances := make(map[uuid.UUID]float64)

	for rows.Next() {
		user, dist, err := scanUser(rows)
		if err != nil {
			return nil, nil, errors.WrapFail(err, "scan rows")
		}
		users = append(users, user)
		distances[user.ID] = dist / 1000
	}

	if err := rows.Err(); err != nil {
		return nil, nil, errors.WrapFail(err, "rows error")
	}

	return users, distances, nil
}

func scanUser(row pgx.Row) (*user.User, float64, error) {
	var u user.User
	var distance sql.NullFloat64
	err := row.Scan(
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
		&distance,
	)
	if err != nil {
		return nil, 0, errors.WrapFail(err, "scan user")
	}
	if !distance.Valid {
		return nil, 0, errors.Error("distance is null")
	}
	return &u, distance.Float64, nil
}

func buildSearchQuery(
	filter *search.Filter,
	interests *interest.Interest,
	curLatitude, curLongitude float64,
	offset, limit uint64,
) (string, []any, error) {
	qb := baseQuery().Where(baseConditions(filter))

	qb = applyMainFilters(qb, filter)
	qb = applyInterestFilters(qb, interests)
	qb = applyDistanceFilter(qb, filter, curLatitude, curLongitude)
	qb = applyPagination(qb, offset, limit)

	return qb.ToSql()
}

func baseQuery() sq.SelectBuilder {
	return sq.Select(
		"u.id", "u.name", "u.age", "u.gender", "u.latitude", "u.longitude",
		"u.bio", "u.goal", "u.zodiac", "u.height", "u.education",
		"u.children", "u.alcohol", "u.smoking", "u.is_hidden",
		"u.is_verified", "u.is_premium", "is_blocked", "u.created_at", "u.updated_at",
	).From("users u").
		PlaceholderFormat(sq.Dollar)
}

func baseConditions(filter *search.Filter) sq.Sqlizer {
	return sq.And{
		sq.Eq{"is_hidden": false},
		sq.Eq{"is_blocked": false},
		sq.NotEq{"id": filter.UserID},
		sq.Expr("NOT EXISTS (SELECT 1 FROM user_views WHERE viewer_id = ? AND viewed_id = u.id)", filter.UserID),
	}
}

func applyMainFilters(qb sq.SelectBuilder, filter *search.Filter) sq.SelectBuilder {
	qb = applyRangeFilter(qb, "age", filter.MinAge, filter.MaxAge)
	qb = applyRangeFilter(qb, "height", filter.MinHeight, filter.MaxHeight)

	if filter.GenderPriority != 0 && filter.GenderPriority != 3 { // 3 - ANY
		qb = qb.Where(sq.Eq{"gender": int(filter.GenderPriority)})
	}

	filters := map[string]int{
		interest.GoalDBName:      int(filter.Goal),
		interest.ZodiacDBName:    int(filter.Zodiac),
		interest.EducationDBName: int(filter.Education),
		interest.ChildrenDBName:  int(filter.Children),
		interest.AlcoholDBName:   int(filter.Alcohol),
		interest.SmokingDBName:   int(filter.Smoking),
	}

	for field, value := range filters {
		if value != 0 {
			qb = qb.Where(sq.Eq{field: value})
		}
	}

	return qb
}

func applyDistanceFilter(
	qb sq.SelectBuilder,
	filter *search.Filter,
	curLatitude, curLongitude float64,
) sq.SelectBuilder {
	distanceExpr := "earth_distance(ll_to_earth(?, ?), ll_to_earth(u.latitude, u.longitude))"
	qb = qb.Column(distanceExpr+" AS distance", curLatitude, curLongitude)

	if filter.MinDistanceKM > 0 {
		qb = qb.Where(sq.Expr(distanceExpr+" >= ?", curLatitude, curLongitude, float64(filter.MinDistanceKM)*1000))
	}
	if filter.MaxDistanceKM > 0 {
		qb = qb.Where(sq.Expr(distanceExpr+" <= ?", curLatitude, curLongitude, float64(filter.MaxDistanceKM)*1000))
	}

	return qb.OrderBy("distance ASC")
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

	for _, pair := range filterInterests {
		qb = qb.Where(sq.Expr(
			"EXISTS (SELECT 1 FROM interests i WHERE i.user_id = u.id AND i.type = ? AND i.value = ?)",
			pair.Type,
			pair.Value,
		))
	}

	return qb
}

func extractInterestPairs(interests *interest.Interest) []struct {
	Type  string
	Value int
} {
	interestGroups := interest.MapInterestsToGroups(interests)

	var pairs []struct {
		Type  string
		Value int
	}
	for t, values := range interestGroups {
		for _, value := range values {
			if value != 0 {
				pairs = append(pairs, struct {
					Type  string
					Value int
				}{t, value})
			}
		}
	}
	return pairs
}

func applyPagination(qb sq.SelectBuilder, offset, limit uint64) sq.SelectBuilder {
	if limit > 0 {
		qb = qb.Limit(limit)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	return qb
}
