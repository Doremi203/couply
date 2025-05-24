package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) FetchIncomingLikes(ctx context.Context, in *dto.FetchIncomingLikesV1Request) (*dto.FetchIncomingLikesV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	likes, err := c.matchingStorageFacade.FetchIncomingLikesTx(ctx, userID, in.Limit, in.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "c.matchingStorageFacade.FetchIncomingLikesTx")
	}

	return &dto.FetchIncomingLikesV1Response{
		Likes: likes,
	}, nil
}
