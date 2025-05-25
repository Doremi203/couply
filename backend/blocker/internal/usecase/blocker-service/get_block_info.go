package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) GetBlockInfo(ctx context.Context, _ *dto.GetBlockInfoV1Request) (*dto.GetBlockInfoV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	block, err := c.blockerStorageFacade.GetBlockInfoTx(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "blockerStorageFacade.GetBlockInfoTx")
	}

	return dto.UserBlockToGetBlockInfoResponse(block), nil
}
