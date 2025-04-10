package search

import desc "github.com/Doremi203/Couply/backend/pkg/search-service/v1"

type GenderPriority int

const (
	GenderPriorityUnspecified GenderPriority = iota
	GenderPriorityMale
	GenderPriorityFemale
	GenderPriorityAny
)

func PBToGenderPriority(genderPriority desc.GenderPriority) GenderPriority {
	switch genderPriority {
	case desc.GenderPriority_GENDERPRIORITY_UNSPECIFIED:
		return GenderPriorityUnspecified
	case desc.GenderPriority_GENDERPRIORITY_MALE:
		return GenderPriorityMale
	case desc.GenderPriority_GENDERPRIORITY_FEMALE:
		return GenderPriorityFemale
	case desc.GenderPriority_GENDERPRIORITY_ANY:
		return GenderPriorityAny
	default:
		return GenderPriority(0)
	}
}

func GenderPriorityToPB(genderPriority GenderPriority) desc.GenderPriority {
	switch genderPriority {
	case GenderPriorityUnspecified:
		return desc.GenderPriority_GENDERPRIORITY_UNSPECIFIED
	case GenderPriorityMale:
		return desc.GenderPriority_GENDERPRIORITY_MALE
	case GenderPriorityFemale:
		return desc.GenderPriority_GENDERPRIORITY_FEMALE
	case GenderPriorityAny:
		return desc.GenderPriority_GENDERPRIORITY_ANY
	default:
		return desc.GenderPriority(0)
	}
}
