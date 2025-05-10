package user_service

import (
	"github.com/Doremi203/couply/backend/common/libs/slices"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type CreateUserV1Request struct {
	Name                string
	Age                 int32
	Gender              user.Gender
	Location            string
	Bio                 string
	Goal                common.Goal
	Interest            *interest.Interest
	Zodiac              common.Zodiac
	Height              int32
	Education           common.Education
	Children            common.Children
	Alcohol             common.Alcohol
	Smoking             common.Smoking
	Hidden              bool
	Verified            bool
	PhotoUploadRequests []user.PhotoUploadRequest
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

func (x *CreateUserV1Request) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
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

func (x *CreateUserV1Request) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *CreateUserV1Request) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *CreateUserV1Request) GetPhotoUploadRequests() []user.PhotoUploadRequest {
	if x != nil {
		return x.PhotoUploadRequests
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
		PhotoUploadRequests: slices.Map(req.GetPhotoUploadRequests(), func(from *desc.PhotoUploadRequest) user.PhotoUploadRequest {
			return user.PhotoUploadRequest{
				OrderNumber: from.GetOrderNumber(),
				MimeType:    from.GetMimeType(),
			}
		}),
	}
}

func CreateUserResponseToPB(resp *CreateUserV1Response) *desc.CreateUserV1Response {
	return &desc.CreateUserV1Response{
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
