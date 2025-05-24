package postgres

import (
	"context"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageMatching) CreateLike(ctx context.Context, like *matching.Like) error {
	query, args, err := buildCreateLikeQuery(like)
	if err != nil {
		return errors.Wrapf(err, "buildCreateLikeQuery with %v", errors.Token("like", like))
	}

	if err = executeCreateLikeQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateLikeQuery with %v", errors.Token("like", like))
	}

	return nil
}

func buildCreateLikeQuery(like *matching.Like) (string, []any, error) {
	query, args, err := sq.Insert(likesTableName).
		Columns(likesColumns...).
		Values(like.SenderID, like.ReceiverID, like.Message, like.Status, like.CreatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateLikeQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
