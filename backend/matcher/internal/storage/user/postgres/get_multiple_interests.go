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

type GetMultipleInterestsOptions struct {
	UserIDs []uuid.UUID
}

func (s *PgStorageUser) GetMultipleInterests(ctx context.Context, opts GetMultipleInterestsOptions) (map[uuid.UUID]*interest.Interest, error) {
	if len(opts.UserIDs) == 0 {
		return map[uuid.UUID]*interest.Interest{}, nil
	}

	query, args, err := buildGetMultipleInterestsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetMultipleInterestsQuery with %v", errors.Token("options", opts))
	}

	interestsFromDB, err := executeGetMultipleInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrap(err, "executeGetInterestsQuery")
	}

	interestsMap := make(map[uuid.UUID]*interest.Interest)
	for _, i := range interestsFromDB {
		if _, exists := interestsMap[i.UserID]; !exists {
			interestsMap[i.UserID] = interest.NewInterest()
		}

		if err = mapInterestValue(interestsMap[i.UserID], i); err != nil {
			return nil, errors.Wrap(err, "mapInterestValue")
		}
	}

	return interestsMap, nil
}

func buildGetMultipleInterestsQuery(opts GetMultipleInterestsOptions) (string, []any, error) {
	query, args, err := sq.Select(interestsColumns...).
		From(interestsTableName).
		Where(sq.Eq{userIDColumnName: opts.UserIDs}).
		OrderBy(userIDColumnName, typeColumnName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetMultipleInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]DBInterest, error) {
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
