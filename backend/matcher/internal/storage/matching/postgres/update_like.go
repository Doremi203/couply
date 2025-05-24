package postgres

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageMatching) UpdateLike(ctx context.Context, like *matching.Like) error {
	query, args, err := buildUpdateLikeQuery(like)
	if err != nil {
		return errors.Wrapf(err, "buildUpdateLikeQuery with %v", errors.Token("like", like))
	}

	result, err := executeUpdateLikeQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdateLikeQuery with %v", errors.Token("like", like))
	}

	if result.RowsAffected() == 0 {
		return matching.ErrLikeNotFound
	}

	return nil
}

func buildUpdateLikeQuery(like *matching.Like) (string, []any, error) {
	query, args, err := sq.Update(likesTableName).
		Set(statusColumnName, like.Status).
		Where(sq.Eq{
			senderIDColumnName:   like.SenderID,
			receiverIDColumnName: like.ReceiverID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdateLikeQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
