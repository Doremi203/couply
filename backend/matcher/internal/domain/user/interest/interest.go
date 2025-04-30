package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"

type Interest struct {
	Sport           []Sport
	SelfDevelopment []SelfDevelopment
	Art             []Art
	Social          []Social
	Hobby           []Hobby
	Gastronomy      []Gastronomy
}

func NewInterest() *Interest {
	return &Interest{
		Sport:           make([]Sport, 0),
		SelfDevelopment: make([]SelfDevelopment, 0),
		Art:             make([]Art, 0),
		Social:          make([]Social, 0),
		Hobby:           make([]Hobby, 0),
		Gastronomy:      make([]Gastronomy, 0),
	}
}

func InterestToPB(interest *Interest) *desc.Interest {
	return &desc.Interest{
		Sport:           SportSliceToPB(interest.Sport),
		Selfdevelopment: SelfDevelopmentSliceToPB(interest.SelfDevelopment),
		Art:             ArtSliceToPB(interest.Art),
		Social:          SocialSliceToPB(interest.Social),
		Hobby:           HobbySliceToPB(interest.Hobby),
		Gastronomy:      GastronomySliceToPB(interest.Gastronomy),
	}
}

func PBToInterest(interest *desc.Interest) *Interest {
	return &Interest{
		Sport:           PBToSportSlice(interest.GetSport()),
		SelfDevelopment: PBToSelfDevelopmentSlice(interest.GetSelfdevelopment()),
		Art:             PBToArtSlice(interest.GetArt()),
		Social:          PBToSocialSlice(interest.GetSocial()),
		Hobby:           PBToHobbySlice(interest.GetHobby()),
		Gastronomy:      PBToGastronomySlice(interest.GetGastronomy()),
	}
}
