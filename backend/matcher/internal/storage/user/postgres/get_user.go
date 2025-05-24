package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type GetUserOptions struct {
	UserId uuid.UUID
}

func (s *PgStorageUser) GetUser(ctx context.Context, opts GetUserOptions) (*user.User, error) {
	query, args, err := buildGetUserQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetUserQuery with %v", errors.Token("options", opts))
	}

	u, err := executeGetUserQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetUserQuery with %v", errors.Token("options", opts))
	}

	return u, nil
}

func buildGetUserQuery(opts GetUserOptions) (string, []any, error) {
	query, args, err := sq.Select(usersColumns...).
		From(usersTableName).
		Where(sq.Eq{idColumnName: opts.UserId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetUserQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*user.User, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	pay, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[user.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(user.ErrUserNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return pay, nil
}
