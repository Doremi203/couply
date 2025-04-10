package search

import (
	"github.com/Doremi203/Couply/backend/internal/domain/user"
	"github.com/Doremi203/Couply/backend/internal/domain/user/interest"
	desc "github.com/Doremi203/Couply/backend/pkg/search-service/v1"
)

type Filter struct {
	GenderPriority GenderPriority
	MaxAge         int32
	Distance       int32
	Goal           user.Goal
	SearchTape     *SearchTape
	Interest       *interest.Interest
	Info           *Info
}

func FilterToPB(filter *Filter) *desc.Filter {
	return &desc.Filter{
		GenderPriority: GenderPriorityToPB(filter.GenderPriority),
		MaxAge:         filter.MaxAge,
		Distance:       filter.Distance,
		Goal:           user.GoalToPB(filter.Goal),
		SearchTape:     SearchTapeToPB(filter.SearchTape),
		Interest:       interest.InterestToPB(filter.Interest),
		Info:           InfoToPB(filter.Info),
	}
}

func PBToFilter(filter *desc.Filter) *Filter {
	return &Filter{
		GenderPriority: PBToGenderPriority(filter.GenderPriority),
		MaxAge:         filter.MaxAge,
		Distance:       filter.Distance,
		Goal:           user.PBToGoal(filter.Goal),
		SearchTape:     PBToSearchTape(filter.SearchTape),
		Interest:       interest.PBToInterest(filter.Interest),
		Info:           PBToInfo(filter.Info),
	}
}
