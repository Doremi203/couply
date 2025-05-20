package payment

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	errPaymentNotFound = errors.Error("payment not found")
)

func (s *PgStoragePayment) GetPayment(ctx context.Context, paymentID uuid.UUID) (*payment.Payment, error) {
	query, args, err := sq.Select("id", "status", "updated_at").
		From("payments").
		Where(sq.Eq{
			"id": paymentID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	row := s.txManager.GetQueryEngine(ctx).QueryRow(ctx, query, args...)

	pay := &payment.Payment{}
	err = row.Scan(
		&pay.ID,
		&pay.Status,
		&pay.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errPaymentNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return pay, nil
}
