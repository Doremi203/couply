package payment

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStoragePayment) GetPendingPayments(ctx context.Context) ([]*payment.Payment, error) {
	query, args, err := buildPendingPaymentsQuery()
	if err != nil {
		return nil, errors.Wrap(err, "GetPendingPayments")
	}

	rows, err := executePendingPaymentsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrap(err, "GetPendingPayments")
	}
	defer rows.Close()

	return scanPaymentRows(rows)
}

func buildPendingPaymentsQuery() (string, []any, error) {
	query, args, err := sq.Select(
		"id", "status", "updated_at", "user_id", "subscription_id",
		"amount", "currency", "gateway_id", "created_at",
	).
		From("payments").
		Where(sq.Eq{"status": payment.PaymentStatusPending}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildPendingPaymentsQuery")
	}
	return query, args, nil
}

func executePendingPaymentsQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgx.Rows, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errPaymentsNotFound, "executePendingPaymentsQuery")
		}
		return nil, errors.Wrap(err, "executePendingPaymentsQuery")
	}
	return rows, nil
}

func scanPaymentRows(rows pgx.Rows) ([]*payment.Payment, error) {
	var payments []*payment.Payment

	for rows.Next() {
		pay, err := scanPaymentRow(rows)
		if err != nil {
			return nil, errors.Wrap(err, "scanPaymentRow")
		}
		payments = append(payments, pay)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scanPaymentRows")
	}

	return payments, nil
}

func scanPaymentRow(row pgx.Row) (*payment.Payment, error) {
	pay := &payment.Payment{}
	err := row.Scan(
		&pay.ID,
		&pay.Status,
		&pay.UpdatedAt,
		&pay.UserID,
		&pay.SubscriptionID,
		&pay.Amount,
		&pay.Currency,
		&pay.GatewayID,
		&pay.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "scanPaymentRow")
	}
	return pay, nil
}
