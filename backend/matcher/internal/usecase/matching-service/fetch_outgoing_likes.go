package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) FetchOutgoingLikes(ctx context.Context, in *dto.FetchOutgoingLikesV1Request) (*dto.FetchOutgoingLikesV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	likes, err := c.matchingStorageFacade.FetchOutgoingLikesTx(ctx, userID, in.Limit, in.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "matchingStorageFacade.FetchOutgoingLikesTx")
	}

	return &dto.FetchOutgoingLikesV1Response{
		Likes: likes,
	}, nil
}
