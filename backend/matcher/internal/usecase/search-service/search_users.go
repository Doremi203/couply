package search_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) SearchUsers(ctx context.Context, in *dto.SearchUsersV1Request) (*dto.SearchUsersV1Response, error) {
	users, err := c.searchStorageFacade.SearchUsersTx(ctx, in.UserID, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	return &dto.SearchUsersV1Response{Users: users}, nil
}
