package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type DBInterest struct {
	UserID       uuid.UUID `db:"user_id"`
	InterestType string    `db:"type"`
	Value        int       `db:"value"`
}

type GetInterestsOptions struct {
	UserID uuid.UUID
}

func (s *PgStorageUser) GetInterests(ctx context.Context, opts GetInterestsOptions) (*interest.Interest, error) {
	query, args, err := buildGetInterestsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetInterestsQuery with %v", errors.Token("options", opts))
	}

	interestsFromDB, err := executeGetInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetInterestsQuery with %v", errors.Token("options", opts))
	}

	userInterests := interest.NewInterest()
	for _, i := range interestsFromDB {
		if err = mapInterestValue(userInterests, i); err != nil {
			return nil, errors.Wrapf(err, "mapInterestValue with %v", errors.Token("options", opts))
		}
	}

	return userInterests, nil
}

func buildGetInterestsQuery(opts GetInterestsOptions) (string, []any, error) {
	query, args, err := sq.Select(interestsColumns...).
		From(interestsTableName).
		Where(sq.Eq{userIDColumnName: opts.UserID}).
		OrderBy(typeColumnName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]DBInterest, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	interests, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBInterest])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(interest.ErrInterestsNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return interests, nil
}

func mapInterestValue(i *interest.Interest, dbInterest DBInterest) error {
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
