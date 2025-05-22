package payment

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	errPaymentNotFound = errors.Error("payment not found")
)

func (s *PgStoragePayment) GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error) {
	query, args, err := buildGetPaymentQuery(paymentID)
	if err != nil {
		return nil, errors.Wrapf(err, "GetPaymentByID with %v", errors.Token("payment_id", paymentID))
	}

	pay, err := executeGetPaymentQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "GetPaymentByID with %v", errors.Token("payment_id", paymentID))
	}

	return pay, nil
}

func buildGetPaymentQuery(paymentID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Select(
		"id", "user_id", "subscription_id", "amount", "currency",
		"status", "gateway_id", "created_at", "updated_at",
	).
		From("payments").
		Where(sq.Eq{"id": paymentID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", nil, errors.Wrap(err, "buildGetPaymentQuery")
	}
	return query, args, nil
}

func executeGetPaymentQuery(ctx context.Context, queryEngine postgres.QueryEngine, query string, args []any) (*payment.Payment, error) {
	row := queryEngine.QueryRow(ctx, query, args...)

	pay := &payment.Payment{}
	err := row.Scan(
		&pay.ID,
		&pay.UserID,
		&pay.SubscriptionID,
		&pay.Amount,
		&pay.Currency,
		&pay.Status,
		&pay.GatewayID,
		&pay.CreatedAt,
		&pay.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errPaymentNotFound, "executeGetPaymentQuery")
		}
		return nil, errors.Wrap(err, "executeGetPaymentQuery")
	}

	return pay, nil
}
