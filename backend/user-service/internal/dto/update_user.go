package dto

import (
	"github.com/Doremi203/Couply/backend/internal/domain"
	"github.com/Doremi203/Couply/backend/internal/domain/interest"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"time"
)

type UpdateUserV1Request struct {
	ID        int64
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

type UpdateUserV1Response struct {
	User domain.User
}

func UpdateUserRequestToPB(req *UpdateUserV1Request) *desc.UpdateUserV1Request {
	return &desc.UpdateUserV1Request{
		Id:        req.ID,
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

func PBToUpdateUserRequest(req *desc.UpdateUserV1Request) *UpdateUserV1Request {
	return &UpdateUserV1Request{
		ID:        req.Id,
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

func UpdateUserRequestToUser(req *UpdateUserV1Request) *domain.User {
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
		SetUpdatedAt(time.Now()).
		Build()

	return user
}

func UpdateUserResponseToPB(resp *UpdateUserV1Response) *desc.UpdateUserV1Response {
	return &desc.UpdateUserV1Response{
		User: domain.UserToPB(&resp.User),
	}
}

func PBToUpdateUserResponse(resp *desc.UpdateUserV1Response) *UpdateUserV1Response {
	return &UpdateUserV1Response{
		User: domain.PBToUser(resp.User),
	}
}
