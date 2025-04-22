package user_service

import (
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user/interest"
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

type UpdateUserV1Request struct {
	ID        int64
	Name      string
	Age       int32
	Gender    user.Gender
	Location  string
	Bio       string
	Goal      user.Goal
	Interest  *interest.Interest
	Zodiac    user.Zodiac
	Height    int32
	Education user.Education
	Children  user.Children
	Alcohol   user.Alcohol
	Smoking   user.Smoking
	Hidden    bool
	Verified  bool
	Photos    []*user.Photo
}

type UpdateUserV1Response struct {
	User *user.User
}

func UpdateUserRequestToPB(req *UpdateUserV1Request) *desc.UpdateUserV1Request {
	return &desc.UpdateUserV1Request{
		Id:        req.ID,
		Name:      req.Name,
		Age:       req.Age,
		Gender:    user.GenderToPB(req.Gender),
		Location:  req.Location,
		Bio:       req.Bio,
		Goal:      user.GoalToPB(req.Goal),
		Interest:  interest.InterestToPB(req.Interest),
		Zodiac:    user.ZodiacToPB(req.Zodiac),
		Height:    req.Height,
		Education: user.EducationToPB(req.Education),
		Children:  user.ChildrenToPB(req.Children),
		Alcohol:   user.AlcoholToPB(req.Alcohol),
		Smoking:   user.SmokingToPB(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    user.PhotoSliceToPB(req.Photos),
	}
}

func PBToUpdateUserRequest(req *desc.UpdateUserV1Request) *UpdateUserV1Request {
	return &UpdateUserV1Request{
		ID:        req.Id,
		Name:      req.Name,
		Age:       req.Age,
		Gender:    user.PBToGender(req.Gender),
		Location:  req.Location,
		Bio:       req.Bio,
		Goal:      user.PBToGoal(req.Goal),
		Interest:  interest.PBToInterest(req.Interest),
		Zodiac:    user.PBToZodiac(req.Zodiac),
		Height:    req.Height,
		Education: user.PBToEducation(req.Education),
		Children:  user.PBToChildren(req.Children),
		Alcohol:   user.PBToAlcohol(req.Alcohol),
		Smoking:   user.PBToSmoking(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    user.PBToPhotoSlice(req.Photos),
	}
}

func UpdateUserRequestToUser(req *UpdateUserV1Request) *user.User {
	return user.NewUserBuilder().
		SetID(req.ID).
		SetName(req.Name).
		SetAge(req.Age).
		SetGender(req.Gender).
		SetLocation(req.Location).
		SetBIO(req.Bio).
		SetGoal(req.Goal).
		SetInterest(req.Interest).
		SetZodiac(req.Zodiac).
		SetHeight(req.Height).
		SetEducation(req.Education).
		SetChildren(req.Children).
		SetAlcohol(req.Alcohol).
		SetSmoking(req.Smoking).
		SetHidden(req.Hidden).
		SetVerified(req.Verified).
		SetPhotos(req.Photos).
		SetUpdatedAt(time.Now()).
		Build()
}

func UpdateUserResponseToPB(resp *UpdateUserV1Response) *desc.UpdateUserV1Response {
	return &desc.UpdateUserV1Response{
		User: user.UserToPB(resp.User),
	}
}

func PBToUpdateUserResponse(resp *desc.UpdateUserV1Response) *UpdateUserV1Response {
	return &UpdateUserV1Response{
		User: user.PBToUser(resp.User),
	}
}
