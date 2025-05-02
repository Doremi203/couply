package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Sport int

const (
	SportUnspecified Sport = iota
	SportRunning
	SportSwimming
	SportYoga
	SportBicycle
	SportGym
	SportSkiing
	SportSnowboarding
	SportDancing
	SportMartialArts
	SportSurfing
	SportHiking
	SportTennis
	SportClimbing
)

func PBToSport(sport desc.Sport) Sport {
	switch sport {
	case desc.Sport_SPORT_UNSPECIFIED:
		return SportUnspecified
	case desc.Sport_SPORT_RUNNING:
		return SportRunning
	case desc.Sport_SPORT_SWIMMING:
		return SportSwimming
	case desc.Sport_SPORT_YOGA:
		return SportYoga
	case desc.Sport_SPORT_BYCICLE:
		return SportBicycle
	case desc.Sport_SPORT_GYM:
		return SportGym
	case desc.Sport_SPORT_SKIING:
		return SportSkiing
	case desc.Sport_SPORT_SNOWBOARDING:
		return SportSnowboarding
	case desc.Sport_SPORT_DANCING:
		return SportDancing
	case desc.Sport_SPORT_MARTIAL_ARTS:
		return SportMartialArts
	case desc.Sport_SPORT_SURFING:
		return SportSurfing
	case desc.Sport_SPORT_HIKING:
		return SportHiking
	case desc.Sport_SPORT_TENNIS:
		return SportTennis
	case desc.Sport_SPORT_CLIMBING:
		return SportClimbing
	default:
		return Sport(0)
	}
}

func SportToPB(sport Sport) desc.Sport {
	switch sport {
	case SportUnspecified:
		return desc.Sport_SPORT_UNSPECIFIED
	case SportRunning:
		return desc.Sport_SPORT_RUNNING
	case SportSwimming:
		return desc.Sport_SPORT_SWIMMING
	case SportYoga:
		return desc.Sport_SPORT_YOGA
	case SportBicycle:
		return desc.Sport_SPORT_BYCICLE
	case SportGym:
		return desc.Sport_SPORT_GYM
	case SportSkiing:
		return desc.Sport_SPORT_SKIING
	case SportSnowboarding:
		return desc.Sport_SPORT_SNOWBOARDING
	case SportDancing:
		return desc.Sport_SPORT_DANCING
	case SportMartialArts:
		return desc.Sport_SPORT_MARTIAL_ARTS
	case SportSurfing:
		return desc.Sport_SPORT_SURFING
	case SportHiking:
		return desc.Sport_SPORT_HIKING
	case SportTennis:
		return desc.Sport_SPORT_TENNIS
	case SportClimbing:
		return desc.Sport_SPORT_CLIMBING
	default:
		return desc.Sport(0)
	}
}

func SportSliceToPB(sports []Sport) []desc.Sport {
	sportsPB := make([]desc.Sport, 0, len(sports))

	for _, sport := range sports {
		sportsPB = append(sportsPB, SportToPB(sport))
	}

	return sportsPB
}

func PBToSportSlice(sports []desc.Sport) []Sport {
	sportsDomain := make([]Sport, 0, len(sports))

	for _, sport := range sports {
		sportsDomain = append(sportsDomain, PBToSport(sport))
	}

	return sportsDomain
}
