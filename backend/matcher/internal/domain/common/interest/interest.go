package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

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

func InterestToPB(interest *Interest) *desc.Interest {
	return &desc.Interest{
		Sport:           SportSliceToPB(interest.GetSport()),
		SelfDevelopment: SelfDevelopmentSliceToPB(interest.GetSelfDevelopment()),
		Art:             ArtSliceToPB(interest.GetArt()),
		Social:          SocialSliceToPB(interest.GetSocial()),
		Hobby:           HobbySliceToPB(interest.GetHobby()),
		Gastronomy:      GastronomySliceToPB(interest.GetGastronomy()),
	}
}

func PBToInterest(interest *desc.Interest) *Interest {
	return &Interest{
		Sport:           PBToSportSlice(interest.GetSport()),
		SelfDevelopment: PBToSelfDevelopmentSlice(interest.GetSelfDevelopment()),
		Art:             PBToArtSlice(interest.GetArt()),
		Social:          PBToSocialSlice(interest.GetSocial()),
		Hobby:           PBToHobbySlice(interest.GetHobby()),
		Gastronomy:      PBToGastronomySlice(interest.GetGastronomy()),
	}
}
