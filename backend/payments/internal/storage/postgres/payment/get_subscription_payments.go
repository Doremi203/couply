package payment

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/jackc/pgx/v5"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var (
	errPaymentsNotFound = errors.Error("payments not found")
)

func (s *PgStoragePayment) GetSubscriptionPayments(ctx context.Context, subID uuid.UUID) ([]string, error) {
	query, args, err := buildSubscriptionPaymentsQuery(subID)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"GetSubscriptionPayments with %v",
			errors.Token("subscription_id", subID),
		)
	}

	rows, err := executeSubscriptionPaymentsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"GetSubscriptionPayments with %v",
			errors.Token("subscription_id", subID),
		)
	}
	defer rows.Close()

	return scanPaymentIDs(rows)
}

func buildSubscriptionPaymentsQuery(subID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Select("id").
		From("payments").
		Where(sq.Eq{"subscription_id": subID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildSubscriptionPaymentsQuery")
	}
	return query, args, nil
}

func executeSubscriptionPaymentsQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgx.Rows, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errPaymentsNotFound, "executeSubscriptionPaymentsQuery")
		}
		return nil, errors.Wrap(err, "executeSubscriptionPaymentsQuery")
	}
	return rows, nil
}

func scanPaymentIDs(rows pgx.Rows) ([]string, error) {
	var ids []string

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "scanPaymentIDs")
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scanPaymentIDs")
	}

	return ids, nil
}
