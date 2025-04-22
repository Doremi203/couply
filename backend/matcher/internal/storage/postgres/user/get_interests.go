package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user/interest"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

func (s *PgStorageUser) GetInterests(ctx context.Context, userID int64) (*interest.Interest, error) {
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

	return i, nil
}

func (s *PgStorageUser) queryInterests(ctx context.Context, userID int64) (pgx.Rows, error) {
	interestsSQL := `
		SELECT type, value 
		FROM Interests 
		WHERE user_id = $1
	`

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, interestsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("GetInterests: %w", err)
	}
	return rows, nil
}

func (s *PgStorageUser) processInterestRow(rows pgx.Rows, i *interest.Interest) error {
	var interestType, value string
	if err := rows.Scan(&interestType, &value); err != nil {
		return fmt.Errorf("GetInterests: %w", err)
	}

	interestType = strings.ToLower(interestType)
	s.parseInterestValue(i, interestType, value)
	return nil
}

func (s *PgStorageUser) parseInterestValue(i *interest.Interest, interestType, value string) {
	switch interestType {
	case "sport":
		addSportInterest(i, value)
	case "self_development":
		addSelfDevelopmentInterest(i, value)
	case "art":
		addArtInterest(i, value)
	case "social":
		addSocialInterest(i, value)
	case "hobby":
		addHobbyInterest(i, value)
	case "gastronomy":
		addGastronomyInterest(i, value)
	default:
		log.Printf("GetInterests: unknown interest type: %s", interestType)
	}
}

// Обработчики для каждого типа интересов
func addSportInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.Sport = append(i.Sport, interest.Sport(c-'0'))
	}
}

func addSelfDevelopmentInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.SelfDevelopment = append(i.SelfDevelopment, interest.SelfDevelopment(c-'0'))
	}
}

func addArtInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.Art = append(i.Art, interest.Art(c-'0'))
	}
}

func addSocialInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.Social = append(i.Social, interest.Social(c-'0'))
	}
}

func addHobbyInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.Hobby = append(i.Hobby, interest.Hobby(c-'0'))
	}
}

func addGastronomyInterest(i *interest.Interest, value string) {
	for _, c := range value {
		i.Gastronomy = append(i.Gastronomy, interest.Gastronomy(c-'0'))
	}
}
