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

func FilterToPB(filter *Filter) *desc.Filter {
	return &desc.Filter{
		UserId:         filter.UserID,
		GenderPriority: GenderPriorityToPB(filter.GenderPriority),
		AgeRange:       &desc.Range{Min: filter.MinAge, Max: filter.MaxAge},
		HeightRange:    &desc.Range{Min: filter.MinHeight, Max: filter.MaxHeight},
		Distance:       filter.Distance,
		Goal:           common.GoalToPB(filter.Goal),
		Zodiac:         common.ZodiacToPB(filter.Zodiac),
		Education:      common.EducationToPB(filter.Education),
		Children:       common.ChildrenToPB(filter.Children),
		Alcohol:        common.AlcoholToPB(filter.Alcohol),
		Smoking:        common.SmokingToPB(filter.Smoking),
		Interest:       interest.InterestToPB(filter.Interest),
		OnlyVerified:   filter.OnlyVerified,
		OnlyPremium:    filter.OnlyPremium,
		CreatedAt:      timestamppb.New(filter.CreatedAt),
		UpdatedAt:      timestamppb.New(filter.UpdatedAt),
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
