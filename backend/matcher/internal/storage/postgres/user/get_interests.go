package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageUser) GetInterests(ctx context.Context, userID uuid.UUID) (*interest.Interest, error) {
	rows, err := s.queryInterests(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	i := interest.NewInterest()
	for rows.Next() {
		if err = s.processInterestRow(rows, i); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetInterests: rows error: %w", err)
	}

	return i, nil
}

func (s *PgStorageUser) queryInterests(ctx context.Context, userID uuid.UUID) (pgx.Rows, error) {
	interestsSQL := `
        SELECT type, value 
        FROM interests 
        WHERE user_id = $1
        ORDER BY type 
    `

	rows, err := s.txManager.GetQueryEngine(ctx).Query(
		ctx,
		interestsSQL,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("GetInterests query failed: %w", err)
	}
	return rows, nil
}

func (s *PgStorageUser) processInterestRow(rows pgx.Rows, i *interest.Interest) error {
	var (
		interestType string
		value        int
	)

	if err := rows.Scan(&interestType, &value); err != nil {
		return fmt.Errorf("GetInterests scan failed: %w", err)
	}

	return s.mapInterestValue(i, interestType, value)
}

func (s *PgStorageUser) mapInterestValue(i *interest.Interest, interestType string, value int) error {
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
		return fmt.Errorf("unknown interest type: %s", interestType)
	}
	return nil
}
