package search

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func (s *PgStorageSearch) DeleteFilterInterests(ctx context.Context, id int64) error {
	interestsSQL := `
		DELETE FROM filter_interests
		WHERE user_id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		interestsSQL,
		id,
	)
	if err != nil {
		return errors.Wrap(err, "DeleteFilterInterests")
	}

	return nil
}
