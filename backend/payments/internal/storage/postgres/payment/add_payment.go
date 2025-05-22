package payment

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	pgerrors "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

var (
	errDuplicatePayment = errors.Error("payment already exists")
)

func (s *PgStoragePayment) AddPayment(ctx context.Context, payment *payment.Payment) error {
	query, args, err := buildInsertPaymentQuery(payment)
	if err != nil {
		return errors.Wrapf(err, "AddPayment with %v", errors.Token("payment", payment))
	}

	if err = executeInsertQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "AddPayment with %v", errors.Token("payment", payment))
	}

	return nil
}

func buildInsertPaymentQuery(payment *payment.Payment) (string, []any, error) {
	query, args, err := sq.Insert("payments").
		Columns(
			"id", "user_id", "subscription_id", "amount", "currency",
			"status", "gateway_id", "created_at", "updated_at",
		).
		Values(
			payment.GetID(), payment.GetUserID(), payment.GetSubscriptionID(), payment.GetAmount(), payment.GetCurrency(),
			payment.GetStatus(), payment.GetGatewayID(), payment.GetCreatedAt(), payment.GetUpdatedAt(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildInsertPaymentQuery")
	}
	return query, args, nil
}

func executeInsertQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if pgerrors.IsUniqueViolationError(err) {
			return errors.Wrap(
				errDuplicatePayment,
				"executeInsertQuery",
			)
		}
		return errors.Wrap(err, "executeInsertQuery")
	}
	return nil
}
