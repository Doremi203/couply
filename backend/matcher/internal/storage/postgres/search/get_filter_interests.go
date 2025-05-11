package search

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageSearch) GetFilterInterests(ctx context.Context, userID uuid.UUID) (*interest.Interest, error) {
	query, args, err := sq.Select("type", "value").
		From("filter_interests").
		Where(sq.Eq{"user_id": userID}).
		OrderBy("type").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	i := interest.NewInterest()
	for rows.Next() {
		var (
			interestType string
			value        int
		)

		if err := rows.Scan(&interestType, &value); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if err := s.mapInterestValue(i, interestType, value); err != nil {
			return nil, fmt.Errorf("failed to map interest: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return i, nil
}

func (s *PgStorageSearch) mapInterestValue(i *interest.Interest, interestType string, value int) error {
	switch interestType {
	case interest.SportDBName:
		i.Sport = append(i.Sport, interest.Sport(value))
	case interest.SelfDevelopmentDBName:
		i.SelfDevelopment = append(i.SelfDevelopment, interest.SelfDevelopment(value))
	case interest.HobbyDBName:
		i.Hobby = append(i.Hobby, interest.Hobby(value))
	case interest.MusicDBName:
		i.Music = append(i.Music, interest.Music(value))
	case interest.MoviesTVDBName:
		i.MoviesTV = append(i.MoviesTV, interest.MoviesTV(value))
	case interest.FoodDrinkDBName:
		i.FoodDrink = append(i.FoodDrink, interest.FoodDrink(value))
	case interest.PersonalityTraitsDBName:
		i.PersonalityTraits = append(i.PersonalityTraits, interest.PersonalityTraits(value))
	case interest.PetsDBName:
		i.Pets = append(i.Pets, interest.Pets(value))
	default:
		return fmt.Errorf("unknown interest type: %s", interestType)
	}
	return nil
}
