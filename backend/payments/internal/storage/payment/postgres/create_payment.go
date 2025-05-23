package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	pgerrors "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

var (
	errDuplicatePayment = errors.Error("payment already exists")
)

func (s *PgStoragePayment) CreatePayment(ctx context.Context, payment *payment.Payment) error {
	query, args, err := buildCreatePaymentQuery(payment)
	if err != nil {
		return errors.Wrapf(
			err,
			"buildCreatePaymentQuery with %v",
			errors.Token("payment", payment),
		)
	}

	if err = executeCreatePaymentQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreatePaymentQuery with %v", errors.Token("payment", payment))
	}

	return nil
}

func buildCreatePaymentQuery(payment *payment.Payment) (string, []any, error) {
	query, args, err := sq.Insert(paymentsTableName).
		Columns(paymentsColumns...).
		Values(payment.ID, payment.UserID, payment.SubscriptionID, payment.Amount, payment.Currency,
			payment.Status, payment.GatewayID, payment.CreatedAt, payment.UpdatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreatePaymentQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if pgerrors.IsUniqueViolationError(err) {
			return errors.Wrap(
				errDuplicatePayment,
				"exec",
			)
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
