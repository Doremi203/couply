package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) GetBlockInfoTx(ctx context.Context, userID uuid.UUID) (*blocker.UserBlock, error) {
	var (
		userBlock *blocker.UserBlock
		reasons   []blocker.ReportReason
		err       error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		userBlock, err = f.storage.GetUserBlock(ctxTx, postgres.GetUserBlockOptions{
			UserID:        userID,
			AcceptedBlock: true,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUserBlock")
		}

		reasons, err = f.storage.GetUserBlockReasons(ctxTx, postgres.GetUserBlockReasonsOptions{
			BlockID: userBlock.ID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUserBlockReasons")
		}

		userBlock.Reasons = reasons

		return nil
	})

	return userBlock, err
}
