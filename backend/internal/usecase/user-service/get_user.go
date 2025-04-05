package user_service

import (
	"context"
	user_service "github.com/Doremi203/Couply/backend/internal/dto/user-service"
)

func (c *UseCase) GetUser(ctx context.Context, in *user_service.GetUserV1Request) (*user_service.GetUserV1Response, error) {
	user, err := c.userStorageFacade.GetUserTx(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &user_service.GetUserV1Response{User: user}, nil
}
