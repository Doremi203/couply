package user_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	photos, err := c.createPhotos(ctx, userID, in.PhotoUploadRequests)
	if err != nil {
		return nil, errors.WrapFail(err, "create photos")
	}

	userToCreate := user.NewUserBuilder().
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
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	if err = c.userStorageFacade.CreateUserTx(ctx, userToCreate); err != nil {
		return nil, err
	}

	return &dto.CreateUserV1Response{User: userToCreate}, nil
}
