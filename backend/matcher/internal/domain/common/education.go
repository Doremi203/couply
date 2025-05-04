package common

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Education int

const (
	EducationUnspecified Education = iota
	EducationSecondary
	EducationHigher
	EducationPHD
)

func PBToEducation(education desc.Education) Education {
	switch education {
	case desc.Education_EDUCATION_UNSPECIFIED:
		return EducationUnspecified
	case desc.Education_EDUCATION_SECONDARY:
		return EducationSecondary
	case desc.Education_EDUCATION_HIGHER:
		return EducationHigher
	case desc.Education_EDUCATION_PHD:
		return EducationPHD
	default:
		return Education(0)
	}
}

func EducationToPB(education Education) desc.Education {
	switch education {
	case EducationUnspecified:
		return desc.Education_EDUCATION_UNSPECIFIED
	case EducationSecondary:
		return desc.Education_EDUCATION_SECONDARY
	case EducationHigher:
		return desc.Education_EDUCATION_HIGHER
	case EducationPHD:
		return desc.Education_EDUCATION_PHD
	default:
		return desc.Education(0)
	}
}
