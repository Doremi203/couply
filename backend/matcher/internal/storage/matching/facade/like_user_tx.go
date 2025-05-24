package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) LikeUserTx(ctx context.Context, like *matching.Like) (*matching.Like, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.AddLike(ctx, like)
		return err
	})

	return like, err
}
