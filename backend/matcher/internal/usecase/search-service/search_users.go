package search_service

import (
	"context"
	"math"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) SearchUsers(ctx context.Context, in *dto.SearchUsersV1Request) (*dto.SearchUsersV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	users, distances, err := c.searchStorageFacade.SearchUsersTx(ctx, userID, in.Offset, in.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "searchStorageFacade.SearchUsersTx")
	}

	response := make([]*search.UserSearchInfo, 0, len(users))
	for _, u := range users {
		dist, ok := distances[u.ID]
		if !ok {
			dist = 0
		}

		err = u.GenerateDownloadPhotoURLS(ctx, c.photoURLGenerator)
		if err != nil {
			c.logger.Warn(errors.Wrap(err, "GenerateDownloadPhotoURLS"))
		}

		response = append(response, &search.UserSearchInfo{
			User:           u,
			DistanceToUser: int32(math.Round(dist)),
		})
	}

	return &dto.SearchUsersV1Response{UsersSearchInfo: response}, nil
}
