package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

type GetFilterInterestsOptions struct {
	UserID uuid.UUID
}

type DBFilterInterest struct {
	UserID       uuid.UUID `db:"user_id"`
	InterestType string    `db:"type"`
	Value        int       `db:"value"`
}

func (s *PgStorageSearch) GetFilterInterests(ctx context.Context, opts GetFilterInterestsOptions) (*interest.Interest, error) {
	query, args, err := buildGetFilterInterestsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetFilterInterestsQuery with %v", errors.Token("options", opts))
	}

	interestsFromDB, err := executeGetFilterInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetFilterInterestsQuery with %v", errors.Token("options", opts))
	}

	filterInterests := interest.NewInterest()
	for _, i := range interestsFromDB {
		if err = mapInterestValue(filterInterests, i); err != nil {
			return nil, errors.Wrapf(err, "mapInterestValue with %v", errors.Token("options", opts))
		}
	}

	return filterInterests, nil
}

func buildGetFilterInterestsQuery(opts GetFilterInterestsOptions) (string, []any, error) {
	query, args, err := sq.Select(filterInterestsColumns...).
		From(filterInterestsTableName).
		Where(sq.Eq{userIdColumnName: opts.UserID}).
		OrderBy(typeColumnName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetFilterInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]DBFilterInterest, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	interests, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBFilterInterest])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}
	if len(interests) == 0 {
		return nil, errors.Wrap(interest.ErrInterestsNotFound, "pgx.CollectRows")
	}

	return interests, nil
}

func mapInterestValue(i *interest.Interest, dbInterest DBFilterInterest) error {
	switch dbInterest.InterestType {
	case interest.SportName:
		i.Sport = append(i.Sport, interest.Sport(dbInterest.Value))
	case interest.SelfDevelopmentName:
		i.SelfDevelopment = append(i.SelfDevelopment, interest.SelfDevelopment(dbInterest.Value))
	case interest.HobbyName:
		i.Hobby = append(i.Hobby, interest.Hobby(dbInterest.Value))
	case interest.MusicName:
		i.Music = append(i.Music, interest.Music(dbInterest.Value))
	case interest.MoviesTVName:
		i.MoviesTV = append(i.MoviesTV, interest.MoviesTV(dbInterest.Value))
	case interest.FoodDrinkName:
		i.FoodDrink = append(i.FoodDrink, interest.FoodDrink(dbInterest.Value))
	case interest.PersonalityTraitsName:
		i.PersonalityTraits = append(i.PersonalityTraits, interest.PersonalityTraits(dbInterest.Value))
	case interest.PetsName:
		i.Pets = append(i.Pets, interest.Pets(dbInterest.Value))
	default:
		return errors.Wrapf(interest.ErrInterestsNotFound, " %v", dbInterest.InterestType)
	}
	return nil
}
