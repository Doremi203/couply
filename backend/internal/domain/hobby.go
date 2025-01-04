package domain

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type Hobby int

const (
	HobbyUnspecified Hobby = iota
	HobbyLiterature
	HobbyVideoGames
	HobbyBoardGames
	HobbyTravels
	HobbyPlantCultivation
	HobbyFishing
	HobbyDogWalks
	HobbyCatsLover
	HobbyCarsAndMotorcycles
	HobbyConcerts
)

func PBToHobby(hobby desc.Hobby) Hobby {
	switch hobby {
	case desc.Hobby_HOBBY_UNSPECIFIED:
		return HobbyUnspecified
	case desc.Hobby_HOBBY_LITERATURE:
		return HobbyLiterature
	case desc.Hobby_HOBBY_VIDEO_GAMES:
		return HobbyVideoGames
	case desc.Hobby_HOBBY_BOARD_GAMES:
		return HobbyBoardGames
	case desc.Hobby_HOBBY_TRAVELS:
		return HobbyTravels
	case desc.Hobby_HOBBY_PLANT_CULTIVATION:
		return HobbyPlantCultivation
	case desc.Hobby_HOBBY_FISHING:
		return HobbyFishing
	case desc.Hobby_HOBBY_DOG_WALKS:
		return HobbyDogWalks
	case desc.Hobby_HOBBY_CATS_LOVER:
		return HobbyCatsLover
	case desc.Hobby_HOBBY_CARS_AND_MOTORCYCLES:
		return HobbyCarsAndMotorcycles
	case desc.Hobby_HOBBY_CONCERTS:
		return HobbyConcerts
	default:
		return Hobby(0)
	}
}

func HobbyToPB(hobby Hobby) desc.Hobby {
	switch hobby {
	case HobbyUnspecified:
		return desc.Hobby_HOBBY_UNSPECIFIED
	case HobbyLiterature:
		return desc.Hobby_HOBBY_LITERATURE
	case HobbyVideoGames:
		return desc.Hobby_HOBBY_VIDEO_GAMES
	case HobbyBoardGames:
		return desc.Hobby_HOBBY_BOARD_GAMES
	case HobbyTravels:
		return desc.Hobby_HOBBY_TRAVELS
	case HobbyPlantCultivation:
		return desc.Hobby_HOBBY_PLANT_CULTIVATION
	case HobbyFishing:
		return desc.Hobby_HOBBY_FISHING
	case HobbyDogWalks:
		return desc.Hobby_HOBBY_DOG_WALKS
	case HobbyCatsLover:
		return desc.Hobby_HOBBY_CATS_LOVER
	case HobbyCarsAndMotorcycles:
		return desc.Hobby_HOBBY_CARS_AND_MOTORCYCLES
	case HobbyConcerts:
		return desc.Hobby_HOBBY_CONCERTS
	default:
		return desc.Hobby(0)
	}
}
