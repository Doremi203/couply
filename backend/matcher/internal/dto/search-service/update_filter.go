package search_service

import (
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type UpdateFilterV1Request struct {
	UserID         int64
	GenderPriority search.GenderPriority
	MinAge         int32
	MaxAge         int32
	MinHeight      int32
	MaxHeight      int32
	Distance       int32
	Goal           common.Goal
	Zodiac         common.Zodiac
	Education      common.Education
	Children       common.Children
	Alcohol        common.Alcohol
	Smoking        common.Smoking
	Interest       *interest.Interest
	OnlyVerified   bool
	OnlyPremium    bool
}

type UpdateFilterV1Response struct {
	Filter *search.Filter
}

func UpdateFilterRequestToPB(req *UpdateFilterV1Request) *desc.UpdateFilterV1Request {
	return &desc.UpdateFilterV1Request{
		UserId:         req.UserID,
		GenderPriority: search.GenderPriorityToPB(req.GenderPriority),
		MinAge:         req.MinAge,
		MaxAge:         req.MaxAge,
		MinHeight:      req.MinHeight,
		MaxHeight:      req.MaxHeight,
		Distance:       req.Distance,
		Goal:           common.GoalToPB(req.Goal),
		Zodiac:         common.ZodiacToPB(req.Zodiac),
		Education:      common.EducationToPB(req.Education),
		Children:       common.ChildrenToPB(req.Children),
		Alcohol:        common.AlcoholToPB(req.Alcohol),
		Smoking:        common.SmokingToPB(req.Smoking),
		Interest:       interest.InterestToPB(req.Interest),
		OnlyVerified:   req.OnlyVerified,
		OnlyPremium:    req.OnlyPremium,
	}
}

func PBToUpdateFilterRequest(req *desc.UpdateFilterV1Request) *UpdateFilterV1Request {
	return &UpdateFilterV1Request{
		UserID:         req.GetUserId(),
		GenderPriority: search.PBToGenderPriority(req.GetGenderPriority()),
		MinAge:         req.GetMinAge(),
		MaxAge:         req.GetMaxAge(),
		MinHeight:      req.GetMinHeight(),
		MaxHeight:      req.GetMaxHeight(),
		Distance:       req.GetDistance(),
		Goal:           common.PBToGoal(req.GetGoal()),
		Zodiac:         common.PBToZodiac(req.GetZodiac()),
		Education:      common.PBToEducation(req.GetEducation()),
		Children:       common.PBToChildren(req.GetChildren()),
		Alcohol:        common.PBToAlcohol(req.GetAlcohol()),
		Smoking:        common.PBToSmoking(req.GetSmoking()),
		Interest:       interest.PBToInterest(req.GetInterest()),
		OnlyVerified:   req.GetOnlyVerified(),
		OnlyPremium:    req.GetOnlyPremium(),
	}
}

func UpdateFilterResponseToPB(resp *UpdateFilterV1Response) *desc.UpdateFilterV1Response {
	return &desc.UpdateFilterV1Response{
		Filter: search.FilterToPB(resp.Filter),
	}
}

func PBToUpdateFilterResponse(resp *desc.UpdateFilterV1Response) *UpdateFilterV1Response {
	return &UpdateFilterV1Response{
		Filter: search.PBToFilter(resp.GetFilter()),
	}
}

func UpdateFilterRequestToFilter(req *UpdateFilterV1Request) *search.Filter {
	return &search.Filter{
		UserID:         req.UserID,
		GenderPriority: req.GenderPriority,
		MinAge:         req.MinAge,
		MaxAge:         req.MaxAge,
		MinHeight:      req.MinHeight,
		MaxHeight:      req.MaxHeight,
		Distance:       req.Distance,
		Goal:           req.Goal,
		Zodiac:         req.Zodiac,
		Education:      req.Education,
		Children:       req.Children,
		Alcohol:        req.Alcohol,
		Smoking:        req.Smoking,
		Interest:       req.Interest,
		OnlyVerified:   req.OnlyVerified,
		OnlyPremium:    req.OnlyPremium,
		UpdatedAt:      time.Now(),
	}
}
