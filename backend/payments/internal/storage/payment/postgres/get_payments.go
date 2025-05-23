package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

var (
	errPaymentsNotFound = errors.Error("payments not found")
)

type GetPaymentsOptions struct {
	PendingPayments bool
	SubscriptionID  uuid.UUID
}

func (s *PgStoragePayment) GetPayments(ctx context.Context, opts GetPaymentsOptions) ([]*payment.Payment, error) {
	query, args, err := buildGetPaymentsQuery(opts)
	if err != nil {
		return nil, errors.Wrap(err, "buildGetPaymentsQuery")
	}

	pays, err := executeGetPaymentsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrap(err, "executeGetPaymentsQuery")
	}

	return pays, nil
}

func buildGetPaymentsQuery(opts GetPaymentsOptions) (string, []any, error) {
	sb := sq.Select(paymentsColumns...).
		From(paymentsTableName)

	if opts.PendingPayments {
		sb = sb.Where(sq.Eq{statusColumn: payment.PaymentStatusPending})
	} else {
		sb = sb.Where(sq.Eq{subscriptionIDColumn: opts.SubscriptionID})
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetPaymentsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]*payment.Payment, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errPaymentsNotFound, "query")
		}
		return nil, errors.Wrap(err, "query")
	}

	pays, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[payment.Payment])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return pays, nil
}
