package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) UpdateMatchTx(ctx context.Context, match *matching.Match) (*matching.Match, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.UpdateMatch(ctx, match)
		return err
	})

	return match, err
}
