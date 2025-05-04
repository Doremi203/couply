package search

import (
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Filter struct {
	UserID         int64
	GenderPriority GenderPriority
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
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (x *Filter) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Filter) GetGenderPriority() GenderPriority {
	if x != nil {
		return x.GenderPriority
	}
	return 0
}

func (x *Filter) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *Filter) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Filter) GetMinHeight() int32 {
	if x != nil {
		return x.MinHeight
	}
	return 0
}

func (x *Filter) GetMaxHeight() int32 {
	if x != nil {
		return x.MaxHeight
	}
	return 0
}

func (x *Filter) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Filter) GetGoal() common.Goal {
	if x != nil {
		return x.Goal
	}
	return 0
}

func (x *Filter) GetZodiac() common.Zodiac {
	if x != nil {
		return x.Zodiac
	}
	return 0
}

func (x *Filter) GetEducation() common.Education {
	if x != nil {
		return x.Education
	}
	return 0
}

func (x *Filter) GetChildren() common.Children {
	if x != nil {
		return x.Children
	}
	return 0
}

func (x *Filter) GetAlcohol() common.Alcohol {
	if x != nil {
		return x.Alcohol
	}
	return 0
}

func (x *Filter) GetSmoking() common.Smoking {
	if x != nil {
		return x.Smoking
	}
	return 0
}

func (x *Filter) GetInterest() *interest.Interest {
	if x != nil {
		return x.Interest
	}
	return nil
}

func (x *Filter) GetOnlyVerified() bool {
	if x != nil {
		return x.OnlyVerified
	}
	return false
}

func (x *Filter) GetOnlyPremium() bool {
	if x != nil {
		return x.OnlyPremium
	}
	return false
}

func (x *Filter) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func (x *Filter) GetUpdatedAt() time.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return time.Time{}
}

func FilterToPB(filter *Filter) *desc.Filter {
	return &desc.Filter{
		UserId:         filter.GetUserID(),
		GenderPriority: GenderPriorityToPB(filter.GetGenderPriority()),
		AgeRange:       &desc.Range{Min: filter.GetMinAge(), Max: filter.GetMaxAge()},
		HeightRange:    &desc.Range{Min: filter.GetMinHeight(), Max: filter.GetMaxHeight()},
		Distance:       filter.GetDistance(),
		Goal:           common.GoalToPB(filter.GetGoal()),
		Zodiac:         common.ZodiacToPB(filter.GetZodiac()),
		Education:      common.EducationToPB(filter.GetEducation()),
		Children:       common.ChildrenToPB(filter.GetChildren()),
		Alcohol:        common.AlcoholToPB(filter.GetAlcohol()),
		Smoking:        common.SmokingToPB(filter.GetSmoking()),
		Interest:       interest.InterestToPB(filter.GetInterest()),
		OnlyVerified:   filter.GetOnlyVerified(),
		OnlyPremium:    filter.GetOnlyPremium(),
		CreatedAt:      timestamppb.New(filter.GetCreatedAt()),
		UpdatedAt:      timestamppb.New(filter.GetUpdatedAt()),
	}
}

func PBToFilter(filter *desc.Filter) *Filter {
	return &Filter{
		UserID:         filter.GetUserId(),
		GenderPriority: PBToGenderPriority(filter.GetGenderPriority()),
		MinAge:         filter.GetAgeRange().GetMin(),
		MaxAge:         filter.GetAgeRange().GetMax(),
		MinHeight:      filter.GetHeightRange().GetMin(),
		MaxHeight:      filter.GetHeightRange().GetMax(),
		Distance:       filter.GetDistance(),
		Goal:           common.PBToGoal(filter.GetGoal()),
		Zodiac:         common.PBToZodiac(filter.GetZodiac()),
		Education:      common.PBToEducation(filter.GetEducation()),
		Children:       common.PBToChildren(filter.GetChildren()),
		Alcohol:        common.PBToAlcohol(filter.GetAlcohol()),
		Smoking:        common.PBToSmoking(filter.GetSmoking()),
		Interest:       interest.PBToInterest(filter.GetInterest()),
		OnlyVerified:   filter.GetOnlyVerified(),
		OnlyPremium:    filter.GetOnlyPremium(),
		CreatedAt:      filter.GetCreatedAt().AsTime(),
		UpdatedAt:      filter.GetUpdatedAt().AsTime(),
	}
}
