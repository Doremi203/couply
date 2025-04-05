package user_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/dto/user-service"
)

func (c *UseCase) DeleteUser(ctx context.Context, in *user_service.DeleteUserV1Request) (*user_service.DeleteUserV1Response, error) {
	err := c.userStorageFacade.DeleteUserTx(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &user_service.DeleteUserV1Response{}, nil
}
