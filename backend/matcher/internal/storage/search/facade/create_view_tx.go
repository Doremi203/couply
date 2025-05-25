package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeSearch) CreateViewTx(ctx context.Context, viewerID, viewedID uuid.UUID) error {
	err := f.searchStorage.CreateUserView(ctx, viewerID, viewedID)
	if err != nil {
		return errors.Wrap(err, "searchStorage.CreateUserView")
	}

	return nil
}
