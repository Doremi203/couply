package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
)

func (f *StorageFacadeBlocker) ReportUserTx(ctx context.Context, block *blocker.UserBlock) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.AddUserBlock(ctxTx, block)
		if err != nil {
			return errors.WrapFail(err, "add user block")
		}

		for _, reason := range block.Reasons {
			if err = f.storage.AddUserBlockReason(ctxTx, block.ID, reason); err != nil {
				return errors.WrapFail(err, "add user block reason")
			}
		}
		return nil
	})

	if err != nil {
		return errors.WrapFail(err, "report user transaction")
	}

	return nil
}
