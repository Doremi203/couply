package user_service

import (
	"context"

	user_service "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) GetUser(ctx context.Context, _ *user_service.GetUserV1Request) (*user_service.GetUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := c.userStorageFacade.GetUserTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &user_service.GetUserV1Response{User: user}, nil
}
