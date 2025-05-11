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

	photos, err := c.createPhotos(ctx, userID, in.GetPhotoUploadRequests())
	if err != nil {
		return nil, errors.WrapFail(err, "create photos")
	}

	user := user.NewUserBuilder().
		SetID(userID).
		SetName(in.GetName()).
		SetAge(in.GetAge()).
		SetGender(in.GetGender()).
		SetLatitude(in.GetLatitude()).
		SetLongitude(in.GetLongitude()).
		SetBIO(in.GetBio()).
		SetGoal(in.GetGoal()).
		SetInterest(in.GetInterest()).
		SetZodiac(in.GetZodiac()).
		SetHeight(in.GetHeight()).
		SetEducation(in.GetEducation()).
		SetChildren(in.GetChildren()).
		SetAlcohol(in.GetAlcohol()).
		SetSmoking(in.GetSmoking()).
		SetIsHidden(in.GetIsHidden()).
		SetIsVerified(in.GetIsVerified()).
		SetIsPremium(in.GetIsPremium()).
		SetIsBlocked(in.GetIsBlocked()).
		SetPhotos(photos).
		SetUpdatedAt(time.Now()).
		Build()

	updatedUser, err := c.userStorageFacade.UpdateUserTx(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserV1Response{User: updatedUser}, nil
}
