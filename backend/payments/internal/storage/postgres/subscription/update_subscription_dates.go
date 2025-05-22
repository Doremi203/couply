package subscription

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/jackc/pgx/v5/pgconn"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageSubscription) UpdateSubscriptionDatesByID(ctx context.Context, subID uuid.UUID, startDate, endDate time.Time) error {
	query, args, err := buildUpdateSubscriptionDatesQuery(subID, startDate, endDate)
	if err != nil {
		return errors.Wrapf(err, "UpdateSubscriptionDatesByID with %v", subID)
	}

	result, err := executeUpdateSubscriptionDatesQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "UpdateSubscriptionDatesByID with %v", subID)
	}

	if err = verifyUpdateResult(result); err != nil {
		return errors.Wrapf(err, "UpdateSubscriptionDatesByID with %v", subID)
	}

	return nil
}

func buildUpdateSubscriptionDatesQuery(subID uuid.UUID, startDate, endDate time.Time) (string, []any, error) {
	query, args, err := sq.Update("subscriptions").
		Set("start_date", startDate).
		Set("end_date", endDate).
		Where(sq.Eq{
			"id": subID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildUpdateSubscriptionDatesQuery")
	}
	return query, args, nil
}

func executeUpdateSubscriptionDatesQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "executeUpdateSubscriptionDatesQuery")
	}
	return result, nil
}

func verifyUpdateResult(result pgconn.CommandTag) error {
	rowsAffected := result.RowsAffected()
	switch {
	case rowsAffected == 0:
		return errors.Wrap(errSubscriptionNotFound, "verifyUpdateResult")
	default:
		return nil
	}
}
