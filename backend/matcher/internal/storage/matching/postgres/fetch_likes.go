package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type FetchLikesOptions struct {
	SenderUserID   uuid.UUID
	ReceiverUserID uuid.UUID
	Incoming       bool
	Outgoing       bool
	Limit          uint64
	Offset         uint64
}

func (s *PgStorageMatching) FetchLikes(ctx context.Context, opts FetchLikesOptions) ([]*matching.Like, error) {
	query, args, err := buildFetchLikesQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildFetchLikesQuery with %v", errors.Token("options", opts))
	}

	likes, err := executeFetchLikesQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeFetchLikesQuery with %v", errors.Token("options", opts))
	}

	return likes, nil
}

func buildFetchLikesQuery(opts FetchLikesOptions) (string, []any, error) {
	sb := sq.Select(likesColumns...).
		From(likesTableName)

	if opts.Incoming {
		sb = sb.Where(sq.Eq{receiverIDColumnName: opts.ReceiverUserID}).
			Where(sq.Eq{statusColumnName: matching.StatusWaiting})
	}
	if opts.Outgoing {
		sb = sb.Where(sq.Eq{senderIDColumnName: opts.SenderUserID}).
			Where(sq.Eq{statusColumnName: matching.StatusWaiting})
	}

	return sb.OrderBy(createdAtColumnName + " DESC").
		Limit(opts.Limit).
		Offset(opts.Offset).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func executeFetchLikesQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]*matching.Like, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	likes, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[matching.Like])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	if len(likes) == 0 {
		return nil, errors.Wrap(matching.ErrLikesNotFound, "pgx.CollectRows")
	}

	return likes, nil
}
