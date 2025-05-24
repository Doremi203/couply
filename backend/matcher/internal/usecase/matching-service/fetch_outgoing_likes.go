package matching_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) FetchOutgoingLikes(ctx context.Context, in *dto.FetchOutgoingLikesV1Request) (*dto.FetchOutgoingLikesV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	likes, err := c.matchingStorageFacade.FetchOutgoingLikesTx(ctx, userID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return &dto.FetchOutgoingLikesV1Response{
		Likes: likes,
	}, nil
}
