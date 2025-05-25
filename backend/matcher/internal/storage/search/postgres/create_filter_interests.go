package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageSearch) CreateFilterInterests(ctx context.Context, userID uuid.UUID, filterInterests *interest.Interest) error {
	interestGroups := interest.MapInterestsToGroups(filterInterests)

	for interestType, values := range interestGroups {
		if len(values) == 0 {
			continue
		}

		if err := s.processFilterInterestGroup(ctx, userID, interestType, values); err != nil {
			return errors.Wrapf(err, "processInterestGroup with %v", errors.Token("user_id", userID))
		}
	}

	return nil
}

func (s *PgStorageSearch) processFilterInterestGroup(ctx context.Context, userID uuid.UUID, interestType string, values []int) error {
	for _, value := range values {
		query, args, err := buildCreateFilterInterestsQuery(userID, interestType, value)
		if err != nil {
			return errors.Wrap(err, "buildCreateFilterInterestsQuery")
		}

		if err = executeCreateFilterInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
			return errors.Wrap(err, "executeCreateFilterInterestsQuery")
		}
	}
	return nil
}

func buildCreateFilterInterestsQuery(userID uuid.UUID, interestType string, value int) (string, []any, error) {
	query, args, err := sq.Insert(filterInterestsTableName).
		Columns(filterInterestsColumns...).
		Values(userID, interestType, value).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateFilterInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsForeignKeyViolationError(err) {
			return user.ErrUserDoesntExist
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
