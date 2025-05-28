package user_service

import (
	"time"

	"github.com/Doremi203/couply/backend/common/libs/slices"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

type UpdateUserV1Request struct {
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
	UpdatedAt           time.Time
	PhotoUploadRequests []user.PhotoUploadRequest
}

type UpdateUserV1Response struct {
	User *user.User
}

func PBToUpdateUserRequest(req *desc.UpdateUserV1Request) *UpdateUserV1Request {
	latitudeWithNoise, longitudeWithNoise := utils.AddNoise(req.GetLatitude(), req.GetLongitude())
	now := time.Now()

	return &UpdateUserV1Request{
		Name:       req.GetName(),
		Age:        req.GetAge(),
		Gender:     user.PBToGender(req.GetGender()),
		Latitude:   latitudeWithNoise,
		Longitude:  longitudeWithNoise,
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
		UpdatedAt:  now,
		PhotoUploadRequests: slices.Map(req.GetPhotoUploadRequests(), func(from *desc.PhotoUploadRequest) user.PhotoUploadRequest {
			return user.PhotoUploadRequest{
				OrderNumber: from.GetOrderNumber(),
				MimeType:    from.GetMimeType(),
			}
		}),
	}
}

func UpdateUserResponseToPB(resp *UpdateUserV1Response) *desc.UpdateUserV1Response {
	return &desc.UpdateUserV1Response{
		User: user.UserToPB(resp.User),
		PhotoUploadResponses: slices.Map(resp.User.Photos, func(from user.Photo) *desc.PhotoUploadResponse {
			var uploadURL string
			if from.UploadURL != nil {
				uploadURL = *from.UploadURL
			}
			return &desc.PhotoUploadResponse{
				OrderNumber: from.OrderNumber,
				UploadUrl:   uploadURL,
			}
		}),
	}
}
