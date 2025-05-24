package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) GetUser(ctx context.Context, _ *dto.GetUserV1Request) (*dto.GetUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContex")
	}

	user, err := c.userStorageFacade.GetUserTx(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "userStorageFacade.GetUser")
	}

	err = user.GenerateDownloadPhotoURLS(ctx, c.photoURLGenerator)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateDownloadPhotoURLS")
	}

	return &dto.GetUserV1Response{User: user}, nil
}
