package interest

type Pets int

const (
	PetsUnspecified Pets = iota
	PetsDogs
	PetsCats
	PetsOther
	PetsNone
)

func PBToPets(p desc.Pets) Pets {
	switch p {
	case desc.Pets_PETS_UNSPECIFIED:
		return PetsUnspecified
	case desc.Pets_PETS_DOGS:
		return PetsDogs
	case desc.Pets_PETS_CATS:
		return PetsCats
	case desc.Pets_PETS_OTHER:
		return PetsOther
	case desc.Pets_PETS_NONE:
		return PetsNone
	default:
		return PetsUnspecified
	}
}

func PetsToPB(p Pets) desc.Pets {
	switch p {
	case PetsUnspecified:
		return desc.Pets_PETS_UNSPECIFIED
	case PetsDogs:
		return desc.Pets_PETS_DOGS
	case PetsCats:
		return desc.Pets_PETS_CATS
	case PetsOther:
		return desc.Pets_PETS_OTHER
	case PetsNone:
		return desc.Pets_PETS_NONE
	default:
		return desc.Pets_PETS_UNSPECIFIED
	}
}

func PetsSliceToPB(p []Pets) []desc.Pets {
	pb := make([]desc.Pets, len(p))
	for i, v := range p {
		pb[i] = PetsToPB(v)
	}
	return pb
}

func PBToPetsSlice(pb []desc.Pets) []Pets {
	p := make([]Pets, len(pb))
	for i, v := range pb {
		p[i] = PBToPets(v)
	}
	return p
}
