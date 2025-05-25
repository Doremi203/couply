package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type GetUsersOptions struct {
	UserIDs []uuid.UUID
}

func (s *PgStorageUser) GetUsers(ctx context.Context, opts GetUsersOptions) ([]*user.User, error) {
	if len(opts.UserIDs) == 0 {
		return []*user.User{}, nil
	}

	query, args, err := buildGetUsersQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetUsersQuery with %v", errors.Token("options", opts))
	}

	users, err := executeGetUsersQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetUsersQuery with %v", errors.Token("options", opts))
	}

	return users, nil
}

func buildGetUsersQuery(opts GetUsersOptions) (string, []any, error) {
	query, args, err := sq.Select(usersColumns...).
		From(usersTableName).
		Where(sq.Eq{idColumnName: opts.UserIDs}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetUsersQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]*user.User, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[user.User])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	if len(users) == 0 {
		return nil, errors.Wrap(user.ErrUserNotFound, "pgx.CollectRows")
	}

	return users, nil
}
