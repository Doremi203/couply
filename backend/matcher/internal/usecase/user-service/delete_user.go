package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) DeleteUser(ctx context.Context, _ *dto.DeleteUserV1Request) (*dto.DeleteUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = c.userStorageFacade.DeleteUserTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteUserV1Response{}, nil
}
