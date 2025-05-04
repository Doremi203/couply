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

func (x *UpdateUserV1Request) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateUserV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserV1Request) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserV1Request) GetGender() user.Gender {
	if x != nil {
		return x.Gender
	}
	return user.Gender(0)
}

func (x *UpdateUserV1Request) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateUserV1Request) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UpdateUserV1Request) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return common.Goal(0)
}

func (x *UpdateUserV1Request) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *UpdateUserV1Request) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return common.Zodiac(0)
}

func (x *UpdateUserV1Request) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UpdateUserV1Request) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return common.Education(0)
}

func (x *UpdateUserV1Request) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return common.Children(0)
}

func (x *UpdateUserV1Request) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return common.Alcohol(0)
}

func (x *UpdateUserV1Request) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return common.Smoking(0)
}

func (x *UpdateUserV1Request) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *UpdateUserV1Request) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *UpdateUserV1Request) GetPhotos() []*user.Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

type UpdateUserV1Response struct {
	User *user.User
}

func (x *UpdateUserV1Response) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func UpdateUserRequestToPB(req *UpdateUserV1Request) *desc.UpdateUserV1Request {
	return &desc.UpdateUserV1Request{
		Id:        req.GetID(),
		Name:      req.GetName(),
		Age:       req.GetAge(),
		Gender:    user.GenderToPB(req.GetGender()),
		Location:  req.GetLocation(),
		Bio:       req.GetBio(),
		Goal:      common.GoalToPB(req.GetGoal()),
		Interest:  interest.InterestToPB(req.GetInterest()),
		Zodiac:    common.ZodiacToPB(req.GetZodiac()),
		Height:    req.GetHeight(),
		Education: common.EducationToPB(req.GetEducation()),
		Children:  common.ChildrenToPB(req.GetChildren()),
		Alcohol:   common.AlcoholToPB(req.GetAlcohol()),
		Smoking:   common.SmokingToPB(req.GetSmoking()),
		Hidden:    req.GetHidden(),
		Verified:  req.GetVerified(),
		Photos:    user.PhotoSliceToPB(req.GetPhotos()),
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
		SetID(req.GetID()).
		SetName(req.GetName()).
		SetAge(req.GetAge()).
		SetGender(req.GetGender()).
		SetLocation(req.GetLocation()).
		SetBIO(req.GetBio()).
		SetGoal(req.GetGoal()).
		SetInterest(req.GetInterest()).
		SetZodiac(req.GetZodiac()).
		SetHeight(req.GetHeight()).
		SetEducation(req.GetEducation()).
		SetChildren(req.GetChildren()).
		SetAlcohol(req.GetAlcohol()).
		SetSmoking(req.GetSmoking()).
		SetHidden(req.GetHidden()).
		SetVerified(req.GetVerified()).
		SetPhotos(req.GetPhotos()).
		SetUpdatedAt(time.Now()).
		Build()
}

func UpdateUserResponseToPB(resp *UpdateUserV1Response) *desc.UpdateUserV1Response {
	return &desc.UpdateUserV1Response{
		User: user.UserToPB(resp.GetUser()),
	}
}

func PBToUpdateUserResponse(resp *desc.UpdateUserV1Response) *UpdateUserV1Response {
	return &UpdateUserV1Response{
		User: user.PBToUser(resp.GetUser()),
	}
}
