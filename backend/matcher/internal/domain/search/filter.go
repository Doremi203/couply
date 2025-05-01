package search

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user/interest"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Filter struct {
	UserID            int32
	GenderPriority    GenderPriority
	MinAge            int32
	MaxAge            int32
	MinHeight         int32
	MaxHeight         int32
	Distance          int32
	Goal              user.Goal
	SearchPreferences *SearchPreferences
	Zodiac            user.Zodiac
	Education         user.Education
	Children          user.Children
	Alcohol           user.Alcohol
	Smoking           user.Smoking
	Interest          *interest.Interest
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func FilterToPB(filter *Filter) *desc.Filter {
	return &desc.Filter{
		UserId:            filter.UserID,
		GenderPriority:    GenderPriorityToPB(filter.GenderPriority),
		AgeRange:          &desc.Range{Min: filter.MinAge, Max: filter.MaxAge},
		HeightRange:       &desc.Range{Min: filter.MinHeight, Max: filter.MaxHeight},
		Distance:          filter.Distance,
		Goal:              user.GoalToPB(filter.Goal),
		SearchPreferences: SearchPreferencesToPB(filter.SearchPreferences),
		Zodiac:            user.ZodiacToPB(filter.Zodiac),
		Education:         user.EducationToPB(filter.Education),
		Children:          user.ChildrenToPB(filter.Children),
		Alcohol:           user.AlcoholToPB(filter.Alcohol),
		Smoking:           user.SmokingToPB(filter.Smoking),
		Interest:          interest.InterestToPB(filter.Interest),
		CreatedAt:         timestamppb.New(filter.CreatedAt),
		UpdatedAt:         timestamppb.New(filter.UpdatedAt),
	}
}

func PBToFilter(filter *desc.Filter) *Filter {
	return &Filter{
		UserID:            filter.GetUserId(),
		GenderPriority:    PBToGenderPriority(filter.GetGenderPriority()),
		MinAge:            filter.GetAgeRange().GetMin(),
		MaxAge:            filter.GetAgeRange().GetMax(),
		MinHeight:         filter.GetHeightRange().GetMin(),
		MaxHeight:         filter.GetHeightRange().GetMax(),
		Distance:          filter.GetDistance(),
		Goal:              user.PBToGoal(filter.GetGoal()),
		SearchPreferences: PBToSearchPreferences(filter.GetSearchPreferences()),
		Zodiac:            user.PBToZodiac(filter.GetZodiac()),
		Education:         user.PBToEducation(filter.GetEducation()),
		Children:          user.PBToChildren(filter.GetChildren()),
		Alcohol:           user.PBToAlcohol(filter.GetAlcohol()),
		Smoking:           user.PBToSmoking(filter.GetSmoking()),
		Interest:          interest.PBToInterest(filter.GetInterest()),
		CreatedAt:         filter.GetCreatedAt().AsTime(),
		UpdatedAt:         filter.GetUpdatedAt().AsTime(),
	}
}
