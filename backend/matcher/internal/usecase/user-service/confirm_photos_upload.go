package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) ConfirmPhotosUpload(ctx context.Context, orderNumbers []int32) error {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	err = c.userStorageFacade.UpdatePhotosUploadedAtTx(ctx, orderNumbers, userID)
	if err != nil {
		return errors.WrapFailf(
			err,
			"confirm photos upload for user with %v",
			errors.Token("user_id", userID),
		)
	}

	return nil
}
