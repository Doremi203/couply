package user_service

import (
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type UpdateUserV1Request struct {
	ID        int64
	Name      string
	Age       int32
	Gender    user.Gender
	Location  string
	Bio       string
	Goal      common.Goal
	Interest  *interest.Interest
	Zodiac    common.Zodiac
	Height    int32
	Education common.Education
	Children  common.Children
	Alcohol   common.Alcohol
	Smoking   common.Smoking
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
		Goal:      common.GoalToPB(req.Goal),
		Interest:  interest.InterestToPB(req.Interest),
		Zodiac:    common.ZodiacToPB(req.Zodiac),
		Height:    req.Height,
		Education: common.EducationToPB(req.Education),
		Children:  common.ChildrenToPB(req.Children),
		Alcohol:   common.AlcoholToPB(req.Alcohol),
		Smoking:   common.SmokingToPB(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    user.PhotoSliceToPB(req.Photos),
	}
}

func PBToUpdateUserRequest(req *desc.UpdateUserV1Request) *UpdateUserV1Request {
	return &UpdateUserV1Request{
		ID:        req.GetId(),
		Name:      req.GetName(),
		Age:       req.GetAge(),
		Gender:    user.PBToGender(req.GetGender()),
		Location:  req.GetLocation(),
		Bio:       req.GetBio(),
		Goal:      common.PBToGoal(req.GetGoal()),
		Interest:  interest.PBToInterest(req.GetInterest()),
		Zodiac:    common.PBToZodiac(req.GetZodiac()),
		Height:    req.GetHeight(),
		Education: common.PBToEducation(req.GetEducation()),
		Children:  common.PBToChildren(req.GetChildren()),
		Alcohol:   common.PBToAlcohol(req.GetAlcohol()),
		Smoking:   common.PBToSmoking(req.GetSmoking()),
		Hidden:    req.GetHidden(),
		Verified:  req.GetVerified(),
		Photos:    user.PBToPhotoSlice(req.GetPhotos()),
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
		User: user.PBToUser(resp.GetUser()),
	}
}
