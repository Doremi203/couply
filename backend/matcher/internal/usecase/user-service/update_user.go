package user_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/matcher/utils"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
)

func (c *UseCase) UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	photos, err := c.createPhotos(ctx, userID, in.PhotoUploadRequests)
	if err != nil {
		return nil, errors.Wrap(err, "createPhotos")
	}

	latitudeWithNoise, longitudeWithNoise := utils.AddNoise(in.Latitude, in.Longitude)

	user := user.NewUserBuilder().
		SetID(userID).
		SetName(in.Name).
		SetAge(in.Age).
		SetGender(in.Gender).
		SetLatitude(latitudeWithNoise).
		SetLongitude(longitudeWithNoise).
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
		return nil, errors.Wrap(err, "userStorageFacade.UpdateUserTx")
	}

	return &dto.UpdateUserV1Response{User: updatedUser}, nil
}
