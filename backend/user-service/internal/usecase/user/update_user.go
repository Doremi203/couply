package user

import (
	"context"

	"github.com/Doremi203/Couply/backend/internal/dto"
)

func (c *UseCase) UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error) {
	user := dto.UpdateUserRequestToUser(in)

	updatedUser, err := c.userStorageFacade.UpdateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserV1Response{User: *updatedUser}, nil
}
