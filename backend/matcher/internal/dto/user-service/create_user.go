package user_service

import (
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type CreateUserV1Request struct {
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

type CreateUserV1Response struct {
	User *user.User
}

func CreateUserRequestToPB(req *CreateUserV1Request) *desc.CreateUserV1Request {
	return &desc.CreateUserV1Request{
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

func PBToCreateUserRequest(req *desc.CreateUserV1Request) *CreateUserV1Request {
	return &CreateUserV1Request{
		Name:      req.Name,
		Age:       req.Age,
		Gender:    user.PBToGender(req.Gender),
		Location:  req.Location,
		Bio:       req.Bio,
		Goal:      common.PBToGoal(req.Goal),
		Interest:  interest.PBToInterest(req.Interest),
		Zodiac:    common.PBToZodiac(req.Zodiac),
		Height:    req.Height,
		Education: common.PBToEducation(req.Education),
		Children:  common.PBToChildren(req.Children),
		Alcohol:   common.PBToAlcohol(req.Alcohol),
		Smoking:   common.PBToSmoking(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    user.PBToPhotoSlice(req.Photos),
	}
}

func CreateUserRequestToUser(req *CreateUserV1Request) *user.User {
	return user.NewUserBuilder().
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
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()
}

func CreateUserResponseToPB(resp *CreateUserV1Response) *desc.CreateUserV1Response {
	return &desc.CreateUserV1Response{
		User: user.UserToPB(resp.User),
	}
}

func PBToCreateUserResponse(resp *desc.CreateUserV1Response) *CreateUserV1Response {
	return &CreateUserV1Response{
		User: user.PBToUser(resp.User),
	}
}
