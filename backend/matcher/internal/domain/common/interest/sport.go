package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Sport int

const (
	SportUnspecified Sport = iota
	SportGym
	SportRunning
	SportYoga
	SportSwimming
	SportCycling
	SportTennis
	SportBasketball
	SportHiking
	SportDancing
	SportMartialArts
	SportFootball
	SportSkiing
	SportClimbing
)

func PBToSport(s desc.Sport) Sport {
	switch s {
	case desc.Sport_SPORT_UNSPECIFIED:
		return SportUnspecified
	case desc.Sport_SPORT_GYM:
		return SportGym
	case desc.Sport_SPORT_RUNNING:
		return SportRunning
	case desc.Sport_SPORT_YOGA:
		return SportYoga
	case desc.Sport_SPORT_SWIMMING:
		return SportSwimming
	case desc.Sport_SPORT_CYCLING:
		return SportCycling
	case desc.Sport_SPORT_TENNIS:
		return SportTennis
	case desc.Sport_SPORT_BASKETBALL:
		return SportBasketball
	case desc.Sport_SPORT_HIKING:
		return SportHiking
	case desc.Sport_SPORT_DANCING:
		return SportDancing
	case desc.Sport_SPORT_MARTIAL_ARTS:
		return SportMartialArts
	case desc.Sport_SPORT_FOOTBALL:
		return SportFootball
	case desc.Sport_SPORT_SKIING:
		return SportSkiing
	case desc.Sport_SPORT_CLIMBING:
		return SportClimbing
	default:
		return SportUnspecified
	}
}

func SportToPB(s Sport) desc.Sport {
	switch s {
	case SportUnspecified:
		return desc.Sport_SPORT_UNSPECIFIED
	case SportGym:
		return desc.Sport_SPORT_GYM
	case SportRunning:
		return desc.Sport_SPORT_RUNNING
	case SportYoga:
		return desc.Sport_SPORT_YOGA
	case SportSwimming:
		return desc.Sport_SPORT_SWIMMING
	case SportCycling:
		return desc.Sport_SPORT_CYCLING
	case SportTennis:
		return desc.Sport_SPORT_TENNIS
	case SportBasketball:
		return desc.Sport_SPORT_BASKETBALL
	case SportHiking:
		return desc.Sport_SPORT_HIKING
	case SportDancing:
		return desc.Sport_SPORT_DANCING
	case SportMartialArts:
		return desc.Sport_SPORT_MARTIAL_ARTS
	case SportFootball:
		return desc.Sport_SPORT_FOOTBALL
	case SportSkiing:
		return desc.Sport_SPORT_SKIING
	case SportClimbing:
		return desc.Sport_SPORT_CLIMBING
	default:
		return desc.Sport_SPORT_UNSPECIFIED
	}
}

func SportSliceToPB(s []Sport) []desc.Sport {
	pb := make([]desc.Sport, len(s))
	for i, v := range s {
		pb[i] = SportToPB(v)
	}
	return pb
}

func PBToSportSlice(pb []desc.Sport) []Sport {
	s := make([]Sport, len(pb))
	for i, v := range pb {
		s[i] = PBToSport(v)
	}
	return s
}
