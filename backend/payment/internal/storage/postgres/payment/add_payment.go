package payment

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/payment/internal/domain/payment"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStoragePayment) AddPayment(ctx context.Context, payment *payment.Payment) error {
	query, args, err := sq.Insert("payments").
		Columns("id", "user_id", "subscription_id", "amount", "currency", "status", "gateway_id",
			"created_at", "updated_at").
		Values(
			payment.GetID(),
			payment.GetUserID(),
			payment.GetSubscriptionID(),
			payment.GetAmount(),
			payment.GetCurrency(),
			payment.GetStatus(),
			payment.GetGatewayID(),
			payment.GetCreatedAt(),
			payment.GetUpdatedAt(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
