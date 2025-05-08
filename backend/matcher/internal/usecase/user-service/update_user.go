package user_service

import (
	"context"

	user_service "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) UpdateUser(ctx context.Context, in *user_service.UpdateUserV1Request) (*user_service.UpdateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := user_service.UpdateUserRequestToUser(in, userID)

	updatedUser, err := c.userStorageFacade.UpdateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user_service.UpdateUserV1Response{User: updatedUser}, nil
}
