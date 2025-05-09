package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) GetUsers(ctx context.Context, in *dto.GetUsersV1Request) (*dto.GetUsersV1Response, error) {
	users, err := c.userStorageFacade.GetUsersTx(ctx, in.GetUserIDs())
	if err != nil {
		return nil, err
	}

	return &dto.GetUsersV1Response{Users: users}, nil
}
