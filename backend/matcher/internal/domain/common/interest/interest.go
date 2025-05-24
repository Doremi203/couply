package interest

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"
)

const (
	GoalDBName      = "goal"
	ZodiacDBName    = "zodiac"
	EducationDBName = "education"
	ChildrenDBName  = "children"
	AlcoholDBName   = "alcohol"
	SmokingDBName   = "smoking"

	SportName             = "sport"
	SelfDevelopmentName   = "self_development"
	HobbyName             = "hobby"
	MusicName             = "music"
	MoviesTVName          = "movies_tv"
	FoodDrinkName         = "food_drink"
	PersonalityTraitsName = "personality_traits"
	PetsName              = "pets"
)

var (
	ErrInterestsNotFound   = errors.Error("interest not found")
	ErrUnknownInterestType = errors.Error("unknown interest type")
)

type Interest struct {
	Sport             []Sport
	SelfDevelopment   []SelfDevelopment
	Hobby             []Hobby
	Music             []Music
	MoviesTV          []MoviesTV
	FoodDrink         []FoodDrink
	PersonalityTraits []PersonalityTraits
	Pets              []Pets
}

func NewInterest() *Interest {
	return &Interest{
		Sport:             make([]Sport, 0),
		SelfDevelopment:   make([]SelfDevelopment, 0),
		Hobby:             make([]Hobby, 0),
		Music:             make([]Music, 0),
		MoviesTV:          make([]MoviesTV, 0),
		FoodDrink:         make([]FoodDrink, 0),
		PersonalityTraits: make([]PersonalityTraits, 0),
		Pets:              make([]Pets, 0),
	}
}

func InterestToPB(interest *Interest) *desc.Interest {
	if interest == nil {
		return nil
	}

	return &desc.Interest{
		Sport:             SportSliceToPB(interest.Sport),
		SelfDevelopment:   SelfDevelopmentSliceToPB(interest.SelfDevelopment),
		Hobby:             HobbySliceToPB(interest.Hobby),
		Music:             MusicSliceToPB(interest.Music),
		MoviesTv:          MoviesTVSliceToPB(interest.MoviesTV),
		FoodDrink:         FoodDrinkSliceToPB(interest.FoodDrink),
		PersonalityTraits: PersonalityTraitsSliceToPB(interest.PersonalityTraits),
		Pets:              PetsSliceToPB(interest.Pets),
	}
}

func PBToInterest(pb *desc.Interest) *Interest {
	if pb == nil {
		return nil
	}

	return &Interest{
		Sport:             PBToSportSlice(pb.GetSport()),
		SelfDevelopment:   PBToSelfDevelopmentSlice(pb.GetSelfDevelopment()),
		Hobby:             PBToHobbySlice(pb.GetHobby()),
		Music:             PBToMusicSlice(pb.GetMusic()),
		MoviesTV:          PBToMoviesTVSlice(pb.GetMoviesTv()),
		FoodDrink:         PBToFoodDrinkSlice(pb.GetFoodDrink()),
		PersonalityTraits: PBToPersonalityTraitsSlice(pb.GetPersonalityTraits()),
		Pets:              PBToPetsSlice(pb.GetPets()),
	}
}

func MapInterestsToGroups(interests *Interest) map[string][]int {
	return map[string][]int{
		SportName:             convertSlice(interests.Sport),
		SelfDevelopmentName:   convertSlice(interests.SelfDevelopment),
		HobbyName:             convertSlice(interests.Hobby),
		MusicName:             convertSlice(interests.Music),
		MoviesTVName:          convertSlice(interests.MoviesTV),
		FoodDrinkName:         convertSlice(interests.FoodDrink),
		PersonalityTraitsName: convertSlice(interests.PersonalityTraits),
		PetsName:              convertSlice(interests.Pets),
	}
}

func convertSlice[T ~int](slice []T) []int {
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		result = append(result, int(v))
	}
	return result
}
