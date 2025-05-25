package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageMatching) DeleteMatch(ctx context.Context, userID, targetUserID uuid.UUID) error {
	user1ID, user2ID := orderUserIDs(userID, targetUserID)

	query, args, err := buildDeleteMatchQuery(user1ID, user2ID)
	if err != nil {
		return errors.Wrapf(err, "buildDeleteMatchQuery with %v & %v",
			errors.Token("first_user_id", user1ID),
			errors.Token("second_user_id", user2ID),
		)
	}

	result, err := executeDeleteMatchQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeDeleteMatchQuery with %v & %v",
			errors.Token("first_user_id", user1ID),
			errors.Token("second_user_id", user2ID),
		)
	}

	if result.RowsAffected() == 0 {
		return matching.ErrMatchNotFound
	}

	return nil
}

func buildDeleteMatchQuery(user1ID, user2ID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Delete(matchesTableName).
		Where(sq.Eq{firstUserIDColumnName: user1ID}).
		Where(sq.Eq{secondUserIDColumnName: user2ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeleteMatchQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
