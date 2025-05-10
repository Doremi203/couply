package search_service

import (
	"context"
	"math"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/logger"

	"github.com/Doremi203/couply/backend/matcher/utils"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) SearchUsers(ctx context.Context, in *dto.SearchUsersV1Request) (*dto.SearchUsersV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	users, distances, err := c.searchStorageFacade.SearchUsersTx(
		ctx,
		userID,
		in.Offset,
		in.Limit,
	)
	if err != nil {
		return nil, err
	}

	response := make([]*search.UserSearchInfo, 0, len(users))
	for _, u := range users {
		dist, ok := distances[u.ID]
		if !ok {
			// TODO: решить что делать
			logger.Warnf(ctx, "no distance found for user %s", u.GetID())
		}

		response = append(response, &search.UserSearchInfo{
			User:           u,
			DistanceToUser: int32(math.Round(dist)), // округление по математическим правилам для удобного восприятия
		})
	}

	return &dto.SearchUsersV1Response{UsersSearchInfo: response}, nil
}
