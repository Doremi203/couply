package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (f *StorageFacadeMatching) UpdateLikeTx(ctx context.Context, like *matching.Like) error {
	if err := f.storage.UpdateLike(ctx, like); err != nil {
		return errors.Wrapf(err, "storage.UpdateLike")
	}

	return nil
}
