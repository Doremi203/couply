package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) GetUserByID(ctx context.Context, req *dto.GetUserByIDV1Request) (*dto.GetUserByIDV1Response, error) {
	user, err := c.userStorageFacade.GetUserTx(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}

	err = user.GenerateDownloadPhotoURLS(ctx, c.photoURLGenerator)
	if err != nil {
		return nil, errors.WrapFail(err, "generate download photo urls")
	}

	return &dto.GetUserByIDV1Response{User: user}, nil
}
