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

func (x *UpdateFilterV1Request) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateFilterV1Request) GetGenderPriority() search.GenderPriority {
	if x != nil {
		return x.GenderPriority
	}
	return search.GenderPriority(0)
}

func (x *UpdateFilterV1Request) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *UpdateFilterV1Request) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *UpdateFilterV1Request) GetMinHeight() int32 {
	if x != nil {
		return x.MinHeight
	}
	return 0
}

func (x *UpdateFilterV1Request) GetMaxHeight() int32 {
	if x != nil {
		return x.MaxHeight
	}
	return 0
}

func (x *UpdateFilterV1Request) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *UpdateFilterV1Request) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return common.Goal(0)
}

func (x *UpdateFilterV1Request) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return common.Zodiac(0)
}

func (x *UpdateFilterV1Request) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return common.Education(0)
}

func (x *UpdateFilterV1Request) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return common.Children(0)
}

func (x *UpdateFilterV1Request) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return common.Alcohol(0)
}

func (x *UpdateFilterV1Request) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return common.Smoking(0)
}

func (x *UpdateFilterV1Request) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *UpdateFilterV1Request) GetOnlyVerified() bool {
	if x != nil {
		return x.OnlyVerified
	}
	return false
}

func (x *UpdateFilterV1Request) GetOnlyPremium() bool {
	if x != nil {
		return x.OnlyPremium
	}
	return false
}

type UpdateFilterV1Response struct {
	Filter *search.Filter
}

func (x *UpdateFilterV1Response) GetFilter() *search.Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func UpdateFilterRequestToPB(req *UpdateFilterV1Request) *desc.UpdateFilterV1Request {
	return &desc.UpdateFilterV1Request{
		UserId:         req.GetUserID(),
		GenderPriority: search.GenderPriorityToPB(req.GetGenderPriority()),
		MinAge:         req.GetMinAge(),
		MaxAge:         req.GetMaxAge(),
		MinHeight:      req.GetMinHeight(),
		MaxHeight:      req.GetMaxHeight(),
		Distance:       req.GetDistance(),
		Goal:           common.GoalToPB(req.GetGoal()),
		Zodiac:         common.ZodiacToPB(req.GetZodiac()),
		Education:      common.EducationToPB(req.GetEducation()),
		Children:       common.ChildrenToPB(req.GetChildren()),
		Alcohol:        common.AlcoholToPB(req.GetAlcohol()),
		Smoking:        common.SmokingToPB(req.GetSmoking()),
		Interest:       interest.InterestToPB(req.GetInterest()),
		OnlyVerified:   req.GetOnlyVerified(),
		OnlyPremium:    req.GetOnlyPremium(),
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
		Filter: search.FilterToPB(resp.GetFilter()),
	}
}

func PBToUpdateFilterResponse(resp *desc.UpdateFilterV1Response) *UpdateFilterV1Response {
	return &UpdateFilterV1Response{
		Filter: search.PBToFilter(resp.GetFilter()),
	}
}

func UpdateFilterRequestToFilter(req *UpdateFilterV1Request) *search.Filter {
	return &search.Filter{
		UserID:         req.GetUserID(),
		GenderPriority: req.GetGenderPriority(),
		MinAge:         req.GetMinAge(),
		MaxAge:         req.GetMaxAge(),
		MinHeight:      req.GetMinHeight(),
		MaxHeight:      req.GetMaxHeight(),
		Distance:       req.GetDistance(),
		Goal:           req.GetGoal(),
		Zodiac:         req.GetZodiac(),
		Education:      req.GetEducation(),
		Children:       req.GetChildren(),
		Alcohol:        req.GetAlcohol(),
		Smoking:        req.GetSmoking(),
		Interest:       req.GetInterest(),
		OnlyVerified:   req.GetOnlyVerified(),
		OnlyPremium:    req.GetOnlyPremium(),
		UpdatedAt:      time.Now(),
	}
}
