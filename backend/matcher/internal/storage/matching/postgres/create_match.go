package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageMatching) CreateMatch(ctx context.Context, match *matching.Match) error {
	user1ID, user2ID := orderUserIDs(match.FirstUserID, match.SecondUserID)

	query, args, err := buildCreateMatchQuery(match, user1ID, user2ID)
	if err != nil {
		return errors.Wrapf(err, "buildCreateMatchQuery with %v", errors.Token("match", match))
	}

	if err = executeCreateMatchQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateMatchQuery with %v", errors.Token("match", match))
	}

	return nil
}

func buildCreateMatchQuery(match *matching.Match, user1ID, user2ID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Insert(matchesTableName).
		Columns(matchesColumns...).
		Values(user1ID, user2ID, match.CreatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateMatchQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsForeignKeyViolationError(err) {
			return user.ErrUserDoesntExist
		}
		if postgres.IsUniqueViolationError(err) {
			return matching.ErrMatchAlreadyExists
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
