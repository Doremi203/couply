package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
)

func (f *StorageFacadeBlocker) ReportUserTx(ctx context.Context, block *blocker.UserBlock) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.CreateUserBlock(ctxTx, block)
		if err != nil {
			return errors.Wrap(err, "storage.CreateUserBlock")
		}

		for _, reason := range block.Reasons {
			if err = f.storage.CreateUserBlockReason(ctxTx, block.ID, reason); err != nil {
				return errors.Wrap(err, "storage.CreateUserBlockReason")
			}
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
