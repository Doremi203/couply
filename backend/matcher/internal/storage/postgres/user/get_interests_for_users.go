package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) GetInterestsForUsers(ctx context.Context, userIDs []uuid.UUID) (map[uuid.UUID]*interest.Interest, error) {
	if len(userIDs) == 0 {
		return map[uuid.UUID]*interest.Interest{}, nil
	}

	query, args, err := sq.Select("user_id", "type", "value").
		From("interests").
		Where(sq.Eq{"user_id": userIDs}).
		OrderBy("user_id", "type").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build interests query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute interests query: %w", err)
	}
	defer rows.Close()

	interestsMap := make(map[uuid.UUID]*interest.Interest)
	for rows.Next() {
		var (
			userID       uuid.UUID
			interestType string
			value        int
		)

		err := rows.Scan(&userID, &interestType, &value)
		if err != nil {
			return nil, fmt.Errorf("failed to scan interest row: %w", err)
		}

		if _, exists := interestsMap[userID]; !exists {
			interestsMap[userID] = interest.NewInterest()
		}

		if err := s.mapInterestValue(interestsMap[userID], interestType, value); err != nil {
			return nil, fmt.Errorf("failed to map interest for user %s: %w", userID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("interests rows iteration error: %w", err)
	}

	return interestsMap, nil
}
