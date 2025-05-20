package payment

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStoragePayment) GetPaymentIDsBySubscriptionID(ctx context.Context, subID uuid.UUID) ([]string, error) {
	query, args, err := sq.Select("id").
		From("payments").
		Where(sq.Eq{
			"subscription_id": subID,
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

	var ids []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return ids, nil
}
