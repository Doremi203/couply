package user

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type Smoking int

const (
	SmokingUnspecified Smoking = iota
	SmokingNegatively
	SmokingNeutrally
	SmokingPositively
)

func PBToSmoking(smoking desc.Smoking) Smoking {
	switch smoking {
	case desc.Smoking_SMOKING_UNSPECIFIED:
		return SmokingUnspecified
	case desc.Smoking_SMOKING_NEGATIVELY:
		return SmokingNegatively
	case desc.Smoking_SMOKING_NEUTRALLY:
		return SmokingNeutrally
	case desc.Smoking_SMOKING_POSITIVELY:
		return SmokingPositively
	default:
		return Smoking(0)
	}
}

func SmokingToPB(smoking Smoking) desc.Smoking {
	switch smoking {
	case SmokingUnspecified:
		return desc.Smoking_SMOKING_UNSPECIFIED
	case SmokingNegatively:
		return desc.Smoking_SMOKING_NEGATIVELY
	case SmokingNeutrally:
		return desc.Smoking_SMOKING_NEUTRALLY
	case SmokingPositively:
		return desc.Smoking_SMOKING_POSITIVELY
	default:
		return desc.Smoking(0)
	}
}
