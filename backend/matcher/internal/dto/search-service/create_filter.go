package search_service

import (
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type CreateFilterV1Request struct {
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

func (x *CreateFilterV1Request) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateFilterV1Request) GetGenderPriority() search.GenderPriority {
	if x != nil {
		return x.GenderPriority
	}
	return search.GenderPriority(0)
}

func (x *CreateFilterV1Request) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *CreateFilterV1Request) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *CreateFilterV1Request) GetMinHeight() int32 {
	if x != nil {
		return x.MinHeight
	}
	return 0
}

func (x *CreateFilterV1Request) GetMaxHeight() int32 {
	if x != nil {
		return x.MaxHeight
	}
	return 0
}

func (x *CreateFilterV1Request) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *CreateFilterV1Request) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return 0
}

func (x *CreateFilterV1Request) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return 0
}

func (x *CreateFilterV1Request) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return 0
}

func (x *CreateFilterV1Request) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return 0
}

func (x *CreateFilterV1Request) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return 0
}

func (x *CreateFilterV1Request) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return 0
}

func (x *CreateFilterV1Request) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *CreateFilterV1Request) GetOnlyVerified() bool {
	if x != nil {
		return x.OnlyVerified
	}
	return false
}

func (x *CreateFilterV1Request) GetOnlyPremium() bool {
	if x != nil {
		return x.OnlyPremium
	}
	return false
}

type CreateFilterV1Response struct {
	Filter *search.Filter
}

func (x *CreateFilterV1Response) GetFilter() *search.Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func CreateFilterRequestToPB(req *CreateFilterV1Request) *desc.CreateFilterV1Request {
	return &desc.CreateFilterV1Request{
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

func PBToCreateFilterRequest(req *desc.CreateFilterV1Request) *CreateFilterV1Request {
	return &CreateFilterV1Request{
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

func CreateFilterResponseToPB(resp *CreateFilterV1Response) *desc.CreateFilterV1Response {
	return &desc.CreateFilterV1Response{
		Filter: search.FilterToPB(resp.GetFilter()),
	}
}

func PBToCreateFilterResponse(resp *desc.CreateFilterV1Response) *CreateFilterV1Response {
	return &CreateFilterV1Response{
		Filter: search.PBToFilter(resp.GetFilter()),
	}
}

func CreateFilterRequestToFilter(req *CreateFilterV1Request) *search.Filter {
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
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
