package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"time"
)

type CreateFilterV1Request struct {
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

type CreateFilterV1Response struct {
	Filter *search.Filter
}

func CreateFilterRequestToPB(req *CreateFilterV1Request) *desc.CreateFilterV1Request {
	return &desc.CreateFilterV1Request{
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

func PBToCreateFilterRequest(req *desc.CreateFilterV1Request) *CreateFilterV1Request {
	return &CreateFilterV1Request{
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

func CreateFilterResponseToPB(resp *CreateFilterV1Response) *desc.CreateFilterV1Response {
	return &desc.CreateFilterV1Response{
		Filter: search.FilterToPB(resp.Filter),
	}
}

func PBToCreateFilterResponse(resp *desc.CreateFilterV1Response) *CreateFilterV1Response {
	return &CreateFilterV1Response{
		Filter: search.PBToFilter(resp.GetFilter()),
	}
}

func CreateFilterRequestToFilter(req *CreateFilterV1Request) *search.Filter {
	// TODO: change userID
	return &search.Filter{
		UserID:         0,
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
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
