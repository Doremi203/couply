package search

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/google/uuid"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrFilterNotFound = errors.Error("filter not found")
)

type Filter struct {
	UserID         uuid.UUID
	GenderPriority GenderPriority
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
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FilterToPB(filter *Filter) *desc.Filter {
	return &desc.Filter{
		GenderPriority:  GenderPriorityToPB(filter.GenderPriority),
		AgeRange:        &desc.Range{Min: filter.MinAge, Max: filter.MaxAge},
		HeightRange:     &desc.Range{Min: filter.MinHeight, Max: filter.MaxHeight},
		DistanceKmRange: &desc.Range{Min: filter.MinDistanceKM, Max: filter.MaxDistanceKM},
		Goal:            common.GoalToPB(filter.Goal),
		Zodiac:          common.ZodiacToPB(filter.Zodiac),
		Education:       common.EducationToPB(filter.Education),
		Children:        common.ChildrenToPB(filter.Children),
		Alcohol:         common.AlcoholToPB(filter.Alcohol),
		Smoking:         common.SmokingToPB(filter.Smoking),
		Interest:        interest.InterestToPB(filter.Interest),
		OnlyVerified:    filter.OnlyVerified,
		OnlyPremium:     filter.OnlyPremium,
		CreatedAt:       timestamppb.New(filter.CreatedAt),
		UpdatedAt:       timestamppb.New(filter.UpdatedAt),
	}
}
