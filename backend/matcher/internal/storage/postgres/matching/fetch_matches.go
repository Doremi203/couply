package matching

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageMatching) FetchMatches(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Match, error) {
	query, args, err := sq.Select(
		"first_user_id",
		"second_user_id",
		"created_at",
	).
		From("matches").
		Where(sq.Or{
			sq.Eq{"first_user_id": userID},
			sq.Eq{"second_user_id": userID},
		}).
		OrderBy("created_at DESC").
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("FetchMatches: failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("FetchMatches: %w", err)
	}
	defer rows.Close()

	var matches []*matching.Match
	for rows.Next() {
		var match matching.Match
		err := rows.Scan(
			&match.FirstUserID,
			&match.SecondUserID,
			&match.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("FetchMatches: failed to scan row: %w", err)
		}

		matches = append(matches, &match)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchMatches: rows error: %w", err)
	}

	return matches, nil
}
