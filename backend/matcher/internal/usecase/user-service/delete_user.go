package user_service

import (
	"context"

	user_service "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) DeleteUser(ctx context.Context, in *user_service.DeleteUserV1Request) (*user_service.DeleteUserV1Response, error) {
	err := c.userStorageFacade.DeleteUserTx(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &user_service.DeleteUserV1Response{}, nil
}
