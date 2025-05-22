package payment

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStoragePayment) UpdatePaymentStatus(ctx context.Context, paymentID uuid.UUID, status payment.PaymentStatus) error {
	query, args, err := buildStatusUpdateQuery(paymentID, status)
	if err != nil {
		return errors.Wrapf(err, "UpdatePaymentStatus with %v", errors.Token("payment_id", paymentID))
	}

	result, err := executeStatusUpdate(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "UpdatePaymentStatus with %v", errors.Token("payment_id", paymentID))
	}

	if err = verifyUpdateResult(result); err != nil {
		return errors.Wrapf(err, "UpdatePaymentStatus with %v", errors.Token("payment_id", paymentID))
	}

	return nil
}

func buildStatusUpdateQuery(paymentID uuid.UUID, status payment.PaymentStatus) (string, []any, error) {
	query, args, err := sq.Update("payments").
		Set("status", status).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": paymentID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildStatusUpdateQuery")
	}
	return query, args, nil
}

func executeStatusUpdate(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "executeStatusUpdate")
	}
	return result, nil
}

func verifyUpdateResult(result pgconn.CommandTag) error {
	switch rowsAffected := result.RowsAffected(); rowsAffected {
	case 0:
		return errors.Wrap(errPaymentNotFound, "verifyUpdateResult")
	default:
		return nil
	}
}
