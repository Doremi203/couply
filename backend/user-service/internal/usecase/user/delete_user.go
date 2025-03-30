package user

import (
	"context"

	"github.com/Doremi203/Couply/backend/internal/dto"
)

func (c *UseCase) DeleteUser(ctx context.Context, in *dto.DeleteUserV1Request) (*dto.DeleteUserV1Response, error) {
	err := c.userStorageFacade.DeleteUserTx(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteUserV1Response{}, nil
}
