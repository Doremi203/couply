package payment

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStoragePayment) GetPendingPayments(ctx context.Context) ([]*payment.Payment, error) {
	query, args, err := sq.Select("id", "status", "updated_at", "user_id", "subscription_id", "amount", "currency", "gateway_id", "created_at").
		From("payments").
		Where(sq.Eq{
			"status": payment.PaymentStatusPending,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var payments []*payment.Payment
	for rows.Next() {
		pay := &payment.Payment{}
		err = rows.Scan(
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
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		payments = append(payments, pay)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return payments, nil
}
