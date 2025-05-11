package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

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
		Art:               make([]Art, 0),
		Social:            make([]Social, 0),
		Hobby:             make([]Hobby, 0),
		Gastronomy:        make([]Gastronomy, 0),
		Music:             make([]Music, 0),
		MoviesTV:          make([]MoviesTV, 0),
		FoodDrink:         make([]FoodDrink, 0),
		PersonalityTraits: make([]PersonalityTraits, 0),
		Pets:              make([]Pets, 0),
	}
}

// Getter methods
func (x *Interest) GetSport() []Sport {
	if x != nil {
		return x.Sport
	}
	return nil
}

func (x *Interest) GetSelfDevelopment() []SelfDevelopment {
	if x != nil {
		return x.SelfDevelopment
	}
	return nil
}

func (x *Interest) GetArt() []Art {
	if x != nil {
		return x.Art
	}
	return nil
}

func (x *Interest) GetSocial() []Social {
	if x != nil {
		return x.Social
	}
	return nil
}

func (x *Interest) GetHobby() []Hobby {
	if x != nil {
		return x.Hobby
	}
	return nil
}

func (x *Interest) GetGastronomy() []Gastronomy {
	if x != nil {
		return x.Gastronomy
	}
	return nil
}

func (x *Interest) GetMusic() []Music {
	if x != nil {
		return x.Music
	}
	return nil
}

func (x *Interest) GetMoviesTV() []MoviesTV {
	if x != nil {
		return x.MoviesTV
	}
	return nil
}

func (x *Interest) GetFoodDrink() []FoodDrink {
	if x != nil {
		return x.FoodDrink
	}
	return nil
}

func (x *Interest) GetPersonalityTraits() []PersonalityTraits {
	if x != nil {
		return x.PersonalityTraits
	}
	return nil
}

func (x *Interest) GetPets() []Pets {
	if x != nil {
		return x.Pets
	}
	return nil
}

// Conversion functions
func InterestToPB(interest *Interest) *desc.Interest {
	if interest == nil {
		return nil
	}

	return &desc.Interest{
		Sport:             SportSliceToPB(interest.Sport),
		SelfDevelopment:   SelfDevelopmentSliceToPB(interest.SelfDevelopment),
		Art:               ArtSliceToPB(interest.Art),
		Social:            SocialSliceToPB(interest.Social),
		Hobby:             HobbySliceToPB(interest.Hobby),
		Gastronomy:        GastronomySliceToPB(interest.Gastronomy),
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
		Sport:             PBToSportSlice(pb.Sport),
		SelfDevelopment:   PBToSelfDevelopmentSlice(pb.SelfDevelopment),
		Art:               PBToArtSlice(pb.Art),
		Social:            PBToSocialSlice(pb.Social),
		Hobby:             PBToHobbySlice(pb.Hobby),
		Gastronomy:        PBToGastronomySlice(pb.Gastronomy),
		Music:             PBToMusicSlice(pb.Music),
		MoviesTV:          PBToMoviesTVSlice(pb.MoviesTv),
		FoodDrink:         PBToFoodDrinkSlice(pb.FoodDrink),
		PersonalityTraits: PBToPersonalityTraitsSlice(pb.PersonalityTraits),
		Pets:              PBToPetsSlice(pb.Pets),
	}
}
