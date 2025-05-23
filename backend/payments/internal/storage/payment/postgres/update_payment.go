package postgres

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStoragePayment) UpdatePayment(ctx context.Context, pay *payment.Payment) error {
	query, args, err := buildUpdatePaymentQuery(pay)
	if err != nil {
		return errors.Wrapf(err, "buildUpdatePaymentQuery with %v", errors.Token("payment_id", pay.ID))
	}

	result, err := executeUpdatePayment(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdatePayment with %v", errors.Token("payment_id", pay.ID))
	}

	if err = verifyUpdateResult(result); err != nil {
		return errors.Wrapf(err, "verifyUpdateResult with %v", errors.Token("payment_id", pay.ID))
	}

	return nil
}

func buildUpdatePaymentQuery(pay *payment.Payment) (string, []any, error) {
	query, args, err := sq.Update(paymentsTableName).
		Set(statusColumn, pay.Status).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: pay.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdatePayment(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}

func verifyUpdateResult(result pgconn.CommandTag) error {
	switch rowsAffected := result.RowsAffected(); rowsAffected {
	case 0:
		return errPaymentNotFound
	default:
		return nil
	}
}
