package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageSearch) CreateUserView(ctx context.Context, viewerID, viewedID uuid.UUID) error {
	query, args, err := buildAddUserViewQuery(viewerID, viewedID)
	if err != nil {
		return errors.Wrapf(err, "buildAddUserViewQuery with %v and %v",
			errors.Token("viewer_id", viewerID),
			errors.Token("viewed_id", viewedID))
	}

	if err = executeAddUserViewQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeAddUserViewQuery with %v and %v",
			errors.Token("viewer_id", viewerID),
			errors.Token("viewed_id", viewedID))
	}

	return nil
}

func buildAddUserViewQuery(viewerID, viewedID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Insert(userViewsColumnName).
		Columns(viewerIDColumnName, viewedIDColumnName).
		Values(viewerID, viewedID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}

	query += " ON CONFLICT (viewer_id, viewed_id) DO NOTHING"

	return query, args, nil
}

func executeAddUserViewQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsForeignKeyViolationError(err) {
			return user.ErrUserDoesntExist
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
