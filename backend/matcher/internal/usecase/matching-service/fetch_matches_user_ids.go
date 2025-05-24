package matching_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
	"github.com/google/uuid"
)

func (c *UseCase) FetchMatchesUserIDs(ctx context.Context, in *dto.FetchMatchesUserIDsV1Request) (*dto.FetchMatchesUserIDsV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	matches, err := c.matchingStorageFacade.FetchMatchesTx(ctx, userID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	otherUserIDS := make([]*uuid.UUID, 0)
	for _, match := range matches {
		otherUserID := match.FirstUserID
		if userID == otherUserID {
			otherUserID = match.SecondUserID
		}

		otherUserIDS = append(otherUserIDS, &otherUserID)
	}

	return &dto.FetchMatchesUserIDsV1Response{
		UserIDs: otherUserIDS,
	}, nil
}
