package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetLikeOptions struct {
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	IsWaiting  bool
}

func (s *PgStorageMatching) GetLike(ctx context.Context, opts GetLikeOptions) (*matching.Like, error) {
	query, args, err := buildGetLikeQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetLikeQuery with %v", errors.Token("options", opts))
	}

	like, err := executeGetLikeQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetLikeQuery with %v", errors.Token("options", opts))
	}

	return like, nil
}

func buildGetLikeQuery(opts GetLikeOptions) (string, []any, error) {
	sb := sq.Select(likesColumns...).
		From(likesTableName).
		Where(sq.Eq{
			senderIDColumnName:   opts.SenderID,
			receiverIDColumnName: opts.ReceiverID,
		})

	if opts.IsWaiting {
		sb = sb.Where(sq.Eq{statusColumnName: matching.StatusWaiting})
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetLikeQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*matching.Like, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	like, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[matching.Like])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(matching.ErrLikeNotFound, "pgx.CollectExactlyOneRow")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return like, nil
}
