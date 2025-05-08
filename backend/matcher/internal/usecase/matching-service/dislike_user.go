package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) DislikeUser(ctx context.Context, in *dto.DislikeUserV1Request) (*dto.DislikeUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	updatedLike := &matching.Like{
		SenderID:   in.GetTargetUserID(),
		ReceiverID: userID,
		Status:     matching.StatusDeclined,
	}

	_, err = c.matchingStorageFacade.UpdateLikeTx(ctx, updatedLike)
	if err != nil {
		return nil, err
	}

	return &dto.DislikeUserV1Response{}, nil
}
