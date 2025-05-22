package updater

import (
	"context"
	"github.com/google/uuid"
)

func (u *Updater) updateUserPremiumStatus(ctx context.Context, userID uuid.UUID, isPremium bool) error {
	user, err := u.userClient.GetUserByIDV1(ctx, userID.String())
	if err != nil {
		return err
	}

	user.IsPremium = isPremium
	return u.userClient.UpdateUserByIDV1(ctx, user)
}
