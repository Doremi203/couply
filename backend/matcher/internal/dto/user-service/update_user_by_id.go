package user_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"time"

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
	UpdatedAt           time.Time
	PhotoUploadRequests []user.PhotoUploadRequest
}

type UpdateUserByIDV1Response struct {
	User *user.User
}

func PBToUpdateUserByIDRequest(req *desc.UpdateUserByIDV1Request) (*UpdateUserByIDV1Request, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	now := time.Now()
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
		UpdatedAt:  now,
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
