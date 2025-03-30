package dto

import (
	"time"

	"github.com/Doremi203/Couply/backend/internal/domain"
	"github.com/Doremi203/Couply/backend/internal/domain/interest"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
)

type CreateUserV1Request struct {
	Name      string
	Age       int32
	Gender    domain.Gender
	Location  string
	Bio       string
	Goal      domain.Goal
	Interest  *interest.Interest
	Zodiac    domain.Zodiac
	Height    int32
	Education domain.Education
	Children  domain.Children
	Alcohol   domain.Alcohol
	Smoking   domain.Smoking
	Hidden    bool
	Verified  bool
	Photos    []*domain.Photo
}

type CreateUserV1Response struct {
	User domain.User
}

func CreateUserRequestToPB(req *CreateUserV1Request) *desc.CreateUserV1Request {
	return &desc.CreateUserV1Request{
		Name:      req.Name,
		Age:       req.Age,
		Gender:    domain.GenderToPB(req.Gender),
		Location:  req.Location,
		Bio:       req.Bio,
		Goal:      domain.GoalToPB(req.Goal),
		Interest:  interest.InterestToPB(req.Interest),
		Zodiac:    domain.ZodiacToPB(req.Zodiac),
		Height:    req.Height,
		Education: domain.EducationToPB(req.Education),
		Children:  domain.ChildrenToPB(req.Children),
		Alcohol:   domain.AlcoholToPB(req.Alcohol),
		Smoking:   domain.SmokingToPB(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    domain.PhotoSliceToPB(req.Photos),
	}
}

func PBToCreateUserRequest(req *desc.CreateUserV1Request) *CreateUserV1Request {
	return &CreateUserV1Request{
		Name:      req.Name,
		Age:       req.Age,
		Gender:    domain.PBToGender(req.Gender),
		Location:  req.Location,
		Bio:       req.Bio,
		Goal:      domain.PBToGoal(req.Goal),
		Interest:  interest.PBToInterest(req.Interest),
		Zodiac:    domain.PBToZodiac(req.Zodiac),
		Height:    req.Height,
		Education: domain.PBToEducation(req.Education),
		Children:  domain.PBToChildren(req.Children),
		Alcohol:   domain.PBToAlcohol(req.Alcohol),
		Smoking:   domain.PBToSmoking(req.Smoking),
		Hidden:    req.Hidden,
		Verified:  req.Verified,
		Photos:    domain.PBToPhotoSlice(req.Photos),
	}
}

func CreateUserRequestToUser(req *CreateUserV1Request) *domain.User {
	user := domain.NewUserBuilder().
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

	return user
}

func CreateUserResponseToPB(resp *CreateUserV1Response) *desc.CreateUserV1Response {
	return &desc.CreateUserV1Response{
		User: domain.UserToPB(&resp.User),
	}
}

func PBToCreateUserResponse(resp *desc.CreateUserV1Response) *CreateUserV1Response {
	return &CreateUserV1Response{
		User: domain.PBToUser(resp.User),
	}
}
