package user

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/dto"
)

func (c *UseCase) CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error) {
	user := dto.CreateUserRequestToUser(in)

	createdUser, err := c.userStorageFacade.CreateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserV1Response{User: *createdUser}, nil
}
