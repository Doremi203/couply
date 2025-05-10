package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) GetUsers(ctx context.Context, in *dto.GetUsersV1Request) (*dto.GetUsersV1Response, error) {
	users, err := c.userStorageFacade.GetUsersTx(ctx, in.GetUserIDs())
	if err != nil {
		return nil, err
	}

	for i := range users {
		downloadablePhotos, err := c.downloadablePhotos(ctx, *users[i])
		if err != nil {
			return nil, errors.WrapFail(err, "get downloadable photos")
		}

		users[i].Photos = downloadablePhotos
	}

	return &dto.GetUsersV1Response{Users: users}, nil
}
