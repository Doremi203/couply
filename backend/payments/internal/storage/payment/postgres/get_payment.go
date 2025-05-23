package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/payments/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	errPaymentNotFound = errors.Error("payment not found")
)

type GetPaymentOptions struct {
	PaymentID uuid.UUID
	ForUpdate bool
}

func (s *PgStoragePayment) GetPayment(ctx context.Context, opts GetPaymentOptions) (*payment.Payment, error) {
	query, args, err := buildGetPaymentQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetPaymentQuery with %v", errors.Token("options", opts))
	}

	pay, err := executeGetPaymentQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetPaymentQuery with %v", errors.Token("options", opts))
	}

	return pay, nil
}

func buildGetPaymentQuery(opts GetPaymentOptions) (string, []any, error) {
	sb := sq.Select(paymentsColumns...).
		From(paymentsTableName).
		Where(sq.Eq{idColumn: opts.PaymentID})

	if opts.ForUpdate {
		sb = sb.Suffix("FOR UPDATE")
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetPaymentQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*payment.Payment, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(errPaymentNotFound, "query")
		}
		return nil, errors.Wrap(err, "query")
	}

	pay, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[payment.Payment])
	if err != nil {
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return pay, nil
}
