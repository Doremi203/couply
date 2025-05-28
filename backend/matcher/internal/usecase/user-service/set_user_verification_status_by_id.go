package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

func (c *UseCase) SetUserVerificationStatusByID(
	ctx context.Context,
	id uuid.UUID,
	status user.VerificationStatus,
) error {
	var isVerified bool
	switch status {
	case user.VerificationStatusPass:
		isVerified = true
	default:
		isVerified = false
	}

	err := c.userStorageFacade.UpdateVerificationStatusTx(ctx, id, isVerified)
	if err != nil {
		return errors.Wrap(err, "userStorageFacade.SetUserVerificationStatusByID")
	}

	return nil
}
