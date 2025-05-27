package search_service

import (
	"time"

	"github.com/google/uuid"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type UpdateFilterV1Request struct {
	GenderPriority search.GenderPriority
	MinAge         int32
	MaxAge         int32
	MinHeight      int32
	MaxHeight      int32
	MinDistanceKM  int32
	MaxDistanceKM  int32
	Goal           common.Goal
	Zodiac         common.Zodiac
	Education      common.Education
	Children       common.Children
	Alcohol        common.Alcohol
	Smoking        common.Smoking
	Interest       *interest.Interest
	OnlyVerified   bool
	OnlyPremium    bool
	UpdatedAt      time.Time
}

type UpdateFilterV1Response struct {
	Filter *search.Filter
}

func PBToUpdateFilterRequest(req *desc.UpdateFilterV1Request) *UpdateFilterV1Request {
	return &UpdateFilterV1Request{
		GenderPriority: search.PBToGenderPriority(req.GetGenderPriority()),
		MinAge:         req.GetMinAge(),
		MaxAge:         req.GetMaxAge(),
		MinHeight:      req.GetMinHeight(),
		MaxHeight:      req.GetMaxHeight(),
		MinDistanceKM:  req.GetMinDistanceKm(),
		MaxDistanceKM:  req.GetMaxDistanceKm(),
		Goal:           common.PBToGoal(req.GetGoal()),
		Zodiac:         common.PBToZodiac(req.GetZodiac()),
		Education:      common.PBToEducation(req.GetEducation()),
		Children:       common.PBToChildren(req.GetChildren()),
		Alcohol:        common.PBToAlcohol(req.GetAlcohol()),
		Smoking:        common.PBToSmoking(req.GetSmoking()),
		Interest:       interest.PBToInterest(req.GetInterest()),
		OnlyVerified:   req.GetOnlyVerified(),
		OnlyPremium:    req.GetOnlyPremium(),
		UpdatedAt:      time.Now(),
	}
}

func UpdateFilterResponseToPB(resp *UpdateFilterV1Response) *desc.UpdateFilterV1Response {
	return &desc.UpdateFilterV1Response{
		Filter: search.FilterToPB(resp.Filter),
	}
}

func UpdateFilterRequestToFilter(req *UpdateFilterV1Request, userID uuid.UUID) *search.Filter {
	return &search.Filter{
		UserID:         userID,
		GenderPriority: req.GenderPriority,
		MinAge:         req.MinAge,
		MaxAge:         req.MaxAge,
		MinHeight:      req.MinHeight,
		MaxHeight:      req.MaxHeight,
		MinDistanceKM:  req.MinDistanceKM,
		MaxDistanceKM:  req.MaxDistanceKM,
		Goal:           req.Goal,
		Zodiac:         req.Zodiac,
		Education:      req.Education,
		Children:       req.Children,
		Alcohol:        req.Alcohol,
		Smoking:        req.Smoking,
		Interest:       req.Interest,
		OnlyVerified:   req.OnlyVerified,
		OnlyPremium:    req.OnlyPremium,
		UpdatedAt:      req.UpdatedAt,
	}
}
