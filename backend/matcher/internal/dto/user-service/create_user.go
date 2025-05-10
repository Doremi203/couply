package user_service

import (
	"time"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type CreateUserV1Request struct {
	Name       string
	Age        int32
	Gender     user.Gender
	Latitude   float64
	Longitude  float64
	Bio        string
	Goal       common.Goal
	Interest   *interest.Interest
	Zodiac     common.Zodiac
	Height     int32
	Education  common.Education
	Children   common.Children
	Alcohol    common.Alcohol
	Smoking    common.Smoking
	IsHidden   bool
	IsVerified bool
	IsPremium  bool
	IsBlocked  bool
	Photos     []*user.Photo
}

func (x *CreateUserV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserV1Request) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateUserV1Request) GetGender() user.Gender {
	if x != nil {
		return x.Gender
	}
	return user.Gender(0)
}

func (x *CreateUserV1Request) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreateUserV1Request) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CreateUserV1Request) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *CreateUserV1Request) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return common.Goal(0)
}

func (x *CreateUserV1Request) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *CreateUserV1Request) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return common.Zodiac(0)
}

func (x *CreateUserV1Request) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *CreateUserV1Request) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return common.Education(0)
}

func (x *CreateUserV1Request) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return common.Children(0)
}

func (x *CreateUserV1Request) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return common.Alcohol(0)
}

func (x *CreateUserV1Request) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return common.Smoking(0)
}

func (x *CreateUserV1Request) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *CreateUserV1Request) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *CreateUserV1Request) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

func (x *CreateUserV1Request) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *CreateUserV1Request) GetPhotos() []*user.Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

type CreateUserV1Response struct {
	User *user.User
}

func (x *CreateUserV1Response) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func PBToCreateUserRequest(req *desc.CreateUserV1Request) *CreateUserV1Request {
	return &CreateUserV1Request{
		Name:       req.GetName(),
		Age:        req.GetAge(),
		Gender:     user.PBToGender(req.GetGender()),
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
		Bio:        req.GetBio(),
		Goal:       common.PBToGoal(req.GetGoal()),
		Interest:   interest.PBToInterest(req.GetInterest()),
		Zodiac:     common.PBToZodiac(req.GetZodiac()),
		Height:     req.GetHeight(),
		Education:  common.PBToEducation(req.GetEducation()),
		Children:   common.PBToChildren(req.GetChildren()),
		Alcohol:    common.PBToAlcohol(req.GetAlcohol()),
		Smoking:    common.PBToSmoking(req.GetSmoking()),
		IsHidden:   req.GetIsHidden(),
		IsVerified: req.GetIsVerified(),
		IsPremium:  req.GetIsPremium(),
		IsBlocked:  req.GetIsBlocked(),
		Photos:     user.PBToPhotoSlice(req.GetPhotos()),
	}
}

func CreateUserRequestToUser(req *CreateUserV1Request, id uuid.UUID) *user.User {
	return user.NewUserBuilder().
		SetID(id).
		SetName(req.GetName()).
		SetAge(req.GetAge()).
		SetGender(req.GetGender()).
		SetLatitude(req.GetLatitude()).
		SetLongitude(req.GetLongitude()).
		SetBIO(req.GetBio()).
		SetGoal(req.GetGoal()).
		SetInterest(req.GetInterest()).
		SetZodiac(req.GetZodiac()).
		SetHeight(req.GetHeight()).
		SetEducation(req.GetEducation()).
		SetChildren(req.GetChildren()).
		SetAlcohol(req.GetAlcohol()).
		SetSmoking(req.GetSmoking()).
		SetIsHidden(req.GetIsHidden()).
		SetIsVerified(req.GetIsVerified()).
		SetIsPremium(req.GetIsPremium()).
		SetIsBlocked(req.GetIsBlocked()).
		SetPhotos(req.GetPhotos()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()
}

func CreateUserResponseToPB(resp *CreateUserV1Response) *desc.CreateUserV1Response {
	return &desc.CreateUserV1Response{
		User: user.UserToPB(resp.GetUser()),
	}
}
