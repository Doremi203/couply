package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := dto.UpdateUserRequestToUser(in, userID)

	updatedUser, err := c.userStorageFacade.UpdateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserV1Response{User: updatedUser}, nil
}
