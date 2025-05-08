package matching_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) FetchMatchesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Match, error) {
	var matches []*matching.Match
	var err error

	err = f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		matches, err = f.storage.FetchMatches(ctx, userID, limit, offset)
		return err
	})

	return matches, err
}
