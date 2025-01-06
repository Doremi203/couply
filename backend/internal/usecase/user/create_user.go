package user

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/domain"
	"github.com/Doremi203/Couply/backend/internal/dto"
	"time"
)

func (c *UseCase) CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error) {
	user := domain.NewUserBuilder().
		SetName(in.Name).
		SetAge(in.Age).
		SetGender(in.Gender).
		SetLocation(in.Location).
		SetBIO(in.Bio).
		SetGoal(in.Goal).
		SetInterest(in.Interest).
		SetZodiac(in.Zodiac).
		SetHeight(in.Height).
		SetEducation(in.Education).
		SetChildren(in.Children).
		SetAlcohol(in.Alcohol).
		SetSmoking(in.Smoking).
		SetHidden(in.Hidden).
		SetVerified(in.Verified).
		SetPhotos(in.Photos).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	err := c.userStorageFacade.CreateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserV1Response{User: user}, nil
}
