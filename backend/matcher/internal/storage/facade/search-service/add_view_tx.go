package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeSearch) AddViewTx(ctx context.Context, viewerID, viewedID uuid.UUID) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.searchStorage.AddUserView(ctxTx, viewerID, viewedID)
		return nil
	})

	if err != nil {
		return errors.WrapFail(err, "add view transaction")
	}

	return nil
}
