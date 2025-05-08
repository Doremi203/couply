package user_service

import (
	"context"

	user_service "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) CreateUser(ctx context.Context, in *user_service.CreateUserV1Request) (*user_service.CreateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := user_service.CreateUserRequestToUser(in, userID)

	createdUser, err := c.userStorageFacade.CreateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user_service.CreateUserV1Response{User: createdUser}, nil
}
