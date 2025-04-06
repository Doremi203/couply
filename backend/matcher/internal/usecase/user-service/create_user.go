package user_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/dto/user-service"
)

func (c *UseCase) CreateUser(ctx context.Context, in *user_service.CreateUserV1Request) (*user_service.CreateUserV1Response, error) {
	user := user_service.CreateUserRequestToUser(in)

	createdUser, err := c.userStorageFacade.CreateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user_service.CreateUserV1Response{User: createdUser}, nil
}
