package user_service

import (
	"fmt"

	"github.com/Doremi203/couply/backend/common/libs/slices"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

type UpdateUserByIDV1Request struct {
	ID                  uuid.UUID
	Name                string
	Age                 int32
	Gender              user.Gender
	Latitude            float64
	Longitude           float64
	Bio                 string
	Goal                common.Goal
	Interest            *interest.Interest
	Zodiac              common.Zodiac
	Height              int32
	Education           common.Education
	Children            common.Children
	Alcohol             common.Alcohol
	Smoking             common.Smoking
	IsHidden            bool
	IsVerified          bool
	IsPremium           bool
	IsBlocked           bool
	PhotoUploadRequests []user.PhotoUploadRequest
}

func (x *UpdateUserByIDV1Request) GetID() uuid.UUID {
	if x != nil {
		return x.ID
	}
	return uuid.Nil
}

func (x *UpdateUserByIDV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserByIDV1Request) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserByIDV1Request) GetGender() user.Gender {
	if x != nil {
		return x.Gender
	}
	return user.Gender(0)
}

func (x *UpdateUserByIDV1Request) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateUserByIDV1Request) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UpdateUserByIDV1Request) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UpdateUserByIDV1Request) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return common.Goal(0)
}

func (x *UpdateUserByIDV1Request) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *UpdateUserByIDV1Request) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return common.Zodiac(0)
}

func (x *UpdateUserByIDV1Request) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UpdateUserByIDV1Request) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return common.Education(0)
}

func (x *UpdateUserByIDV1Request) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return common.Children(0)
}

func (x *UpdateUserByIDV1Request) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return common.Alcohol(0)
}

func (x *UpdateUserByIDV1Request) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return common.Smoking(0)
}

func (x *UpdateUserByIDV1Request) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *UpdateUserByIDV1Request) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *UpdateUserByIDV1Request) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

func (x *UpdateUserByIDV1Request) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *UpdateUserByIDV1Request) GetPhotoUploadRequests() []user.PhotoUploadRequest {
	if x != nil {
		return x.PhotoUploadRequests
	}
	return nil
}

type UpdateUserByIDV1Response struct {
	User *user.User
}

func (x *UpdateUserByIDV1Response) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func PBToUpdateUserByIDRequest(req *desc.UpdateUserByIDV1Request) (*UpdateUserByIDV1Request, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %s", req.GetId())
	}
	return &UpdateUserByIDV1Request{
		ID:         id,
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
		PhotoUploadRequests: slices.Map(req.GetPhotoUploadRequests(), func(from *desc.PhotoUploadRequest) user.PhotoUploadRequest {
			return user.PhotoUploadRequest{
				OrderNumber: from.GetOrderNumber(),
				MimeType:    from.GetMimeType(),
			}
		}),
	}, nil
}

func UpdateUserByIDResponseToPB(resp *UpdateUserByIDV1Response) *desc.UpdateUserByIDV1Response {
	return &desc.UpdateUserByIDV1Response{
		User: user.UserToPB(resp.GetUser()),
		PhotoUploadResponses: slices.Map(resp.GetUser().GetPhotos(), func(from user.Photo) *desc.PhotoUploadResponse {
			var uploadURL string
			if from.UploadURL != nil {
				uploadURL = *from.UploadURL
			}
			return &desc.PhotoUploadResponse{
				OrderNumber: from.GetOrderNumber(),
				UploadUrl:   uploadURL,
			}
		}),
	}
}
