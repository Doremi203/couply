package matching_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) FetchIncomingLikes(ctx context.Context, in *dto.FetchIncomingLikesV1Request) (*dto.FetchIncomingLikesV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	likes, err := c.matchingStorageFacade.FetchIncomingLikesTx(ctx, userID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return &dto.FetchIncomingLikesV1Response{
		Likes: likes,
	}, nil
}
