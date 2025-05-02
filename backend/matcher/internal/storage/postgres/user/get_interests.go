package user

import (
	"context"
	"fmt"
	interest2 "github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (s *PgStorageUser) GetInterests(ctx context.Context, userID int64) (*interest2.Interest, error) {
	rows, err := s.queryInterests(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	i := interest2.NewInterest()
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
		FROM interests 
		WHERE user_id = $1
	`

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, interestsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("GetInterests: %w", err)
	}
	return rows, nil
}

func (s *PgStorageUser) processInterestRow(rows pgx.Rows, i *interest2.Interest) error {
	var interestType, value string
	if err := rows.Scan(&interestType, &value); err != nil {
		return fmt.Errorf("GetInterests: %w", err)
	}

	interestType = strings.ToLower(interestType)
	s.parseInterestValue(i, interestType, value)
	return nil
}

func (s *PgStorageUser) parseInterestValue(i *interest2.Interest, interestType, value string) {
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
func addSportInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.Sport = append(i.Sport, interest2.Sport(c-'0'))
	}
}

func addSelfDevelopmentInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.SelfDevelopment = append(i.SelfDevelopment, interest2.SelfDevelopment(c-'0'))
	}
}

func addArtInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.Art = append(i.Art, interest2.Art(c-'0'))
	}
}

func addSocialInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.Social = append(i.Social, interest2.Social(c-'0'))
	}
}

func addHobbyInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.Hobby = append(i.Hobby, interest2.Hobby(c-'0'))
	}
}

func addGastronomyInterest(i *interest2.Interest, value string) {
	for _, c := range value {
		i.Gastronomy = append(i.Gastronomy, interest2.Gastronomy(c-'0'))
	}
}
