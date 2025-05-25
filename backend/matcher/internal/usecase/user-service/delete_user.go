package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) DeleteUser(ctx context.Context, _ *dto.DeleteUserV1Request) (*dto.DeleteUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContex")
	}

	err = c.userStorageFacade.DeleteUserTx(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "userStorageFacade.DeleteUserTx")
	}

	return &dto.DeleteUserV1Response{}, nil
}
