package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) GetUsers(ctx context.Context, in *dto.GetUsersV1Request) (*dto.GetUsersV1Response, error) {
	users, err := c.userStorageFacade.GetUsersTx(ctx, in.UserIDs)
	if err != nil {
		return nil, errors.Wrap(err, "serStorageFacade.GetUsersTx")
	}

	for i := range users {
		err = users[i].GenerateDownloadPhotoURLS(ctx, c.photoURLGenerator)
		if err != nil {
			return nil, errors.Wrap(err, "GenerateDownloadPhotoURLS")
		}
	}

	return &dto.GetUsersV1Response{Users: users}, nil
}
