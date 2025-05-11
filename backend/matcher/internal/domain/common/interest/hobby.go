package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type Hobby int

const (
	HobbyUnspecified Hobby = iota
	HobbyPhotography
	HobbyPainting
	HobbyBoardGames
	HobbyReading
	HobbyCooking
	HobbyGardening
	HobbyTravel
	HobbyWriting
	HobbyChess
	HobbyCrafts
	HobbyAnimals
	HobbyAstrology
)

func PBToHobby(h desc.Hobby) Hobby {
	switch h {
	case desc.Hobby_HOBBY_UNSPECIFIED:
		return HobbyUnspecified
	case desc.Hobby_HOBBY_PHOTOGRAPHY:
		return HobbyPhotography
	case desc.Hobby_HOBBY_PAINTING:
		return HobbyPainting
	case desc.Hobby_HOBBY_BOARDGAMES:
		return HobbyBoardGames
	case desc.Hobby_HOBBY_READING:
		return HobbyReading
	case desc.Hobby_HOBBY_COOKING:
		return HobbyCooking
	case desc.Hobby_HOBBY_GARDENING:
		return HobbyGardening
	case desc.Hobby_HOBBY_TRAVEL:
		return HobbyTravel
	case desc.Hobby_HOBBY_WRITING:
		return HobbyWriting
	case desc.Hobby_HOBBY_CHESS:
		return HobbyChess
	case desc.Hobby_HOBBY_CRAFTS:
		return HobbyCrafts
	case desc.Hobby_HOBBY_ANIMALS:
		return HobbyAnimals
	case desc.Hobby_HOBBY_ASTROLOGY:
		return HobbyAstrology
	default:
		return HobbyUnspecified
	}
}

func HobbyToPB(h Hobby) desc.Hobby {
	switch h {
	case HobbyUnspecified:
		return desc.Hobby_HOBBY_UNSPECIFIED
	case HobbyPhotography:
		return desc.Hobby_HOBBY_PHOTOGRAPHY
	case HobbyPainting:
		return desc.Hobby_HOBBY_PAINTING
	case HobbyBoardGames:
		return desc.Hobby_HOBBY_BOARDGAMES
	case HobbyReading:
		return desc.Hobby_HOBBY_READING
	case HobbyCooking:
		return desc.Hobby_HOBBY_COOKING
	case HobbyGardening:
		return desc.Hobby_HOBBY_GARDENING
	case HobbyTravel:
		return desc.Hobby_HOBBY_TRAVEL
	case HobbyWriting:
		return desc.Hobby_HOBBY_WRITING
	case HobbyChess:
		return desc.Hobby_HOBBY_CHESS
	case HobbyCrafts:
		return desc.Hobby_HOBBY_CRAFTS
	case HobbyAnimals:
		return desc.Hobby_HOBBY_ANIMALS
	case HobbyAstrology:
		return desc.Hobby_HOBBY_ASTROLOGY
	default:
		return desc.Hobby_HOBBY_UNSPECIFIED
	}
}

func HobbySliceToPB(h []Hobby) []desc.Hobby {
	pb := make([]desc.Hobby, len(h))
	for i, v := range h {
		pb[i] = HobbyToPB(v)
	}
	return pb
}

func PBToHobbySlice(pb []desc.Hobby) []Hobby {
	h := make([]Hobby, len(pb))
	for i, v := range pb {
		h[i] = PBToHobby(v)
	}
	return h
}
