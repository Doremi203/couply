package user_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	photos, err := c.createPhotos(ctx, userID, in.PhotoUploadRequests)
	if err != nil {
		return nil, errors.WrapFail(err, "create photos")
	}

	user := user.NewUserBuilder().
		SetID(userID).
		SetName(in.Name).
		SetAge(in.Age).
		SetGender(in.Gender).
		SetLatitude(in.Latitude).
		SetLongitude(in.Longitude).
		SetBIO(in.Bio).
		SetGoal(in.Goal).
		SetInterest(in.Interest).
		SetZodiac(in.Zodiac).
		SetHeight(in.Height).
		SetEducation(in.Education).
		SetChildren(in.Children).
		SetAlcohol(in.Alcohol).
		SetSmoking(in.Smoking).
		SetIsHidden(in.IsHidden).
		SetIsVerified(in.IsVerified).
		SetIsPremium(in.IsPremium).
		SetIsBlocked(in.IsBlocked).
		SetPhotos(photos).
		SetUpdatedAt(time.Now()).
		Build()

	updatedUser, err := c.userStorageFacade.UpdateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserV1Response{User: updatedUser}, nil
}
