package matching_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) DeleteMatch(ctx context.Context, in *dto.DeleteMatchV1Request) (*dto.DeleteMatchV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = c.matchingStorageFacade.DeleteMatchTx(ctx, userID, in.GetTargetUserID())
	if err != nil {
		return nil, err
	}

	return &dto.DeleteMatchV1Response{}, nil
}
