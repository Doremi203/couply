package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func (c *UseCase) ConfirmPhotosUpload(ctx context.Context, orderNumbers []int32) error {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "token.GetUserIDFromContext")
	}

	err = c.userStorageFacade.UpdatePhotosUploadedAtTx(ctx, orderNumbers, userID)
	if err != nil {
		return errors.Wrapf(
			err,
			"userStorageFacade.UpdatePhotosUploadedAtTx with %v",
			errors.Token("user_id", userID),
		)
	}

	return nil
}
