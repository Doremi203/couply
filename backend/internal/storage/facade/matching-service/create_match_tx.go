package matching_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
)

func (f *StorageFacadeMatching) CreateMatchTx(ctx context.Context, match *matching.Match) (*matching.Match, error) {
	err := f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		err := f.storage.AddMatch(ctx, match)
		return err
	})

	return match, err
}
