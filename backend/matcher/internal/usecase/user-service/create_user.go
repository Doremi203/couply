package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := dto.CreateUserRequestToUser(in, userID)

	createdUser, err := c.userStorageFacade.CreateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserV1Response{User: createdUser}, nil
}
