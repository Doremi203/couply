package blocker

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) GetUserBlockReasons(ctx context.Context, blockID uuid.UUID) ([]blocker.ReportReason, error) {
	query, args, err := sq.Select("reason").
		From("user_block_reasons").
		Where(sq.Eq{"block_id": blockID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build reasons query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute reasons query: %w", err)
	}
	defer rows.Close()

	var reasons []blocker.ReportReason
	for rows.Next() {
		var reason blocker.ReportReason
		if err := rows.Scan(&reason); err != nil {
			return nil, fmt.Errorf("failed to scan reason: %w", err)
		}
		reasons = append(reasons, reason)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return reasons, nil
}
