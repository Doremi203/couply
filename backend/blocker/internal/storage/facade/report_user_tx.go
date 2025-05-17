package facade

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) ReportUserTx(ctx context.Context, blockID, blockedUserID uuid.UUID, message string, createdAt time.Time, reasons []blocker.ReportReason) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.AddUserBlock(ctxTx, blockID, blockedUserID, message, createdAt)
		if err != nil {
			return errors.WrapFail(err, "add user block")
		}

		for _, reason := range reasons {
			if err = f.storage.AddUserBlockReason(ctxTx, blockID, reason); err != nil {
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
