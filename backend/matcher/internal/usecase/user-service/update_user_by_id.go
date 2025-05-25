package user_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) UpdateUserByID(ctx context.Context, in *dto.UpdateUserByIDV1Request) (*dto.UpdateUserByIDV1Response, error) {
	photos, err := c.createPhotos(ctx, in.ID, in.PhotoUploadRequests)
	if err != nil {
		return nil, errors.Wrap(err, "createPhotos")
	}

	user := user.NewUserBuilder().
		SetID(in.ID).
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
		return nil, errors.Wrap(err, "userStorageFacade.UpdateUserByID")
	}

	return &dto.UpdateUserByIDV1Response{User: updatedUser}, nil
}
