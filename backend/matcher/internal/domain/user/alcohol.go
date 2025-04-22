package user

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"

type Alcohol int

const (
	AlcoholUnspecified Alcohol = iota
	AlcoholNegatively
	AlcoholNeutrally
	AlcoholPositively
)

func PBToAlcohol(alcohol desc.Alcohol) Alcohol {
	switch alcohol {
	case desc.Alcohol_ALCOHOL_UNSPECIFIED:
		return AlcoholUnspecified
	case desc.Alcohol_ALCOHOL_NEGATIVELY:
		return AlcoholNegatively
	case desc.Alcohol_ALCOHOL_NEUTRALLY:
		return AlcoholNeutrally
	case desc.Alcohol_ALCOHOL_POSITIVELY:
		return AlcoholPositively
	default:
		return Alcohol(0)
	}
}

func AlcoholToPB(alcohol Alcohol) desc.Alcohol {
	switch alcohol {
	case AlcoholUnspecified:
		return desc.Alcohol_ALCOHOL_UNSPECIFIED
	case AlcoholNegatively:
		return desc.Alcohol_ALCOHOL_NEGATIVELY
	case AlcoholNeutrally:
		return desc.Alcohol_ALCOHOL_NEUTRALLY
	case AlcoholPositively:
		return desc.Alcohol_ALCOHOL_POSITIVELY
	default:
		return desc.Alcohol(0)
	}
}
