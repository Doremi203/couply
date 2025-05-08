package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/utils"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) SearchUsers(ctx context.Context, in *dto.SearchUsersV1Request) (*dto.SearchUsersV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	users, err := c.searchStorageFacade.SearchUsersTx(ctx, userID, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	return &dto.SearchUsersV1Response{Users: users}, nil
}
