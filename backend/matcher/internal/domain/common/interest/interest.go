package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

const (
	GoalDBName      = "goal"
	ZodiacDBName    = "zodiac"
	EducationDBName = "education"
	ChildrenDBName  = "children"
	AlcoholDBName   = "alcohol"
	SmokingDBName   = "smoking"

	SportDBName             = "sport"
	SelfDevelopmentDBName   = "self_development"
	HobbyDBName             = "hobby"
	MusicDBName             = "music"
	MoviesTVDBName          = "movies_tv"
	FoodDrinkDBName         = "food_drink"
	PersonalityTraitsDBName = "personality_traits"
	PetsDBName              = "pets"
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
