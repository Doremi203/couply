package search

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageSearch) GetFilterInterests(ctx context.Context, userID int64) (*interest.Interest, error) {
	rows, err := s.queryFilterInterests(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	i := interest.NewInterest()
	for rows.Next() {
		if err = s.processFilterInterestRow(rows, i); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "GetFilterInterests: rows error")
	}

	return i, nil
}

func (s *PgStorageSearch) queryFilterInterests(ctx context.Context, userID int64) (pgx.Rows, error) {
	interestsSQL := `
        SELECT type, value 
        FROM filter_interests 
        WHERE user_id = $1
        ORDER BY type 
    `

	rows, err := s.txManager.GetQueryEngine(ctx).Query(
		ctx,
		interestsSQL,
		userID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "GetFilterInterests query failed")
	}
	return rows, nil
}

func (s *PgStorageSearch) processFilterInterestRow(rows pgx.Rows, i *interest.Interest) error {
	var (
		interestType string
		value        int
	)

	if err := rows.Scan(&interestType, &value); err != nil {
		return errors.WrapFail(err, "GetFilterInterests scan failed")
	}

	return s.mapInterestValue(i, interestType, value)
}

func (s *PgStorageSearch) mapInterestValue(i *interest.Interest, interestType string, value int) error {
	switch interestType {
	case "sport":
		i.Sport = append(i.Sport, interest.Sport(value))
	case "self_development":
		i.SelfDevelopment = append(i.SelfDevelopment, interest.SelfDevelopment(value))
	case "art":
		i.Art = append(i.Art, interest.Art(value))
	case "social":
		i.Social = append(i.Social, interest.Social(value))
	case "hobby":
		i.Hobby = append(i.Hobby, interest.Hobby(value))
	case "gastronomy":
		i.Gastronomy = append(i.Gastronomy, interest.Gastronomy(value))
	default:
		return errors.Errorf("unknown %v", errors.Token("interest_type", interestType))
	}
	return nil
}
