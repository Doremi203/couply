package domain

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type Social int

const (
	SocialUnspecified Social = iota
	SocialVolunteering
	SocialCharity
	SocialEcoActivism
	SocialElderlyCare
	SocialChildcare
	SocialAnimalWelfare
)

func PBToSocial(social desc.Social) Social {
	switch social {
	case desc.Social_SOCIAL_UNSPECIFIED:
		return SocialUnspecified
	case desc.Social_SOCIAL_VOLUNTEERING:
		return SocialVolunteering
	case desc.Social_SOCIAL_CHARITY:
		return SocialCharity
	case desc.Social_SOCIAL_ECO_ACTIVISM:
		return SocialEcoActivism
	case desc.Social_SOCIAL_ELDERLY_CARE:
		return SocialElderlyCare
	case desc.Social_SOCIAL_CHILDCARE:
		return SocialChildcare
	case desc.Social_SOCIAL_ANIMAL_WELFARE:
		return SocialAnimalWelfare
	default:
		return Social(0)
	}
}

func SocialToPB(social Social) desc.Social {
	switch social {
	case SocialUnspecified:
		return desc.Social_SOCIAL_UNSPECIFIED
	case SocialVolunteering:
		return desc.Social_SOCIAL_VOLUNTEERING
	case SocialCharity:
		return desc.Social_SOCIAL_CHARITY
	case SocialEcoActivism:
		return desc.Social_SOCIAL_ECO_ACTIVISM
	case SocialElderlyCare:
		return desc.Social_SOCIAL_ELDERLY_CARE
	case SocialChildcare:
		return desc.Social_SOCIAL_CHILDCARE
	case SocialAnimalWelfare:
		return desc.Social_SOCIAL_ANIMAL_WELFARE
	default:
		return desc.Social(0)
	}
}
