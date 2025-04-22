package user

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

type Gender int

const (
	GenderUnspecified Gender = iota
	GenderMale
	GenderFemale
)

func PBToGender(gender desc.Gender) Gender {
	switch gender {
	case desc.Gender_GENDER_UNSPECIFIED:
		return GenderUnspecified
	case desc.Gender_GENDER_MALE:
		return GenderMale
	case desc.Gender_GENDER_FEMALE:
		return GenderFemale
	default:
		return Gender(0)
	}
}

func GenderToPB(gender Gender) desc.Gender {
	switch gender {
	case GenderUnspecified:
		return desc.Gender_GENDER_UNSPECIFIED
	case GenderMale:
		return desc.Gender_GENDER_MALE
	case GenderFemale:
		return desc.Gender_GENDER_FEMALE
	default:
		return desc.Gender(0)
	}
}
