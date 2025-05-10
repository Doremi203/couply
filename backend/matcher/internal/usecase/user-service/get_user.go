package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) GetUser(ctx context.Context, _ *dto.GetUserV1Request) (*dto.GetUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := c.userStorageFacade.GetUserTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	downloadablePhotos, err := c.downloadablePhotos(ctx, *user)
	if err != nil {
		return nil, errors.WrapFail(err, "get downloadable photos")
	}

	user.Photos = downloadablePhotos

	return &dto.GetUserV1Response{User: user}, nil
}
