package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) FetchOutgoingMatchesTx(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	var matches []*matching.Match
	var err error

	err = f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		matches, err = f.storage.FetchOutgoingMatches(ctx, userID, limit, offset)
		return err
	})

	return matches, err
}
