package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/payments/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStoragePayment) UpdatePaymentStatus(ctx context.Context, paymentID uuid.UUID, newStatus payment.PaymentStatus) error {
	query, args, err := sq.Update("payments").
		Set("status", newStatus).
		Set("updated_at", time.Now()).
		Where(sq.Eq{
			"id": paymentID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	result, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errPaymentNotFound
	}
	if rowsAffected > 1 {
		return fmt.Errorf("unexpected number of rows affected: %d", rowsAffected)
	}

	return nil
}
