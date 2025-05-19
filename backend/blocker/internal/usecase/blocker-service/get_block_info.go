package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"github.com/Doremi203/couply/backend/blocker/utils"
)

func (c *UseCase) GetBlockInfo(ctx context.Context, _ *dto.GetBlockInfoV1Request) (*dto.GetBlockInfoV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	userBlock, err := c.blockerStorageFacade.GetBlockInfoTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.GetBlockInfoV1Response{
		BlockID:       userBlock.GetID(),
		BlockedUserID: userBlock.GetBlockedID(),
		Message:       userBlock.GetMessage(),
		ReportReasons: userBlock.GetReasons(),
		CreatedAt:     userBlock.GetCreatedAt(),
	}, nil
}
