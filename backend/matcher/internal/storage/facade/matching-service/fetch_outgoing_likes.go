package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) FetchOutgoingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error) {
	var likes []*matching.Like
	var err error

	err = f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		likes, err = f.storage.FetchOutgoingLikes(ctx, userID, limit, offset)
		return err
	})

	return likes, err
}
