package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) LikeUserTx(ctx context.Context, like *matching.Like) error {
	if err := f.storage.CreateLike(ctx, like); err != nil {
		return errors.Wrap(err, "storage.CreateLike")
	}

	return nil
}
