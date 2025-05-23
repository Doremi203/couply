package facade

import (
	"context"
	"fmt"

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
		userBlock, err = f.storage.GetUserBlockByUserID(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("failed to get user block: %w", err)
		}

		reasons, err = f.storage.GetUserBlockReasons(ctxTx, userBlock.ID)
		if err != nil {
			return fmt.Errorf("failed to get block reasons: %w", err)
		}

		userBlock.Reasons = reasons
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	return userBlock, nil
}
