package interest

type PersonalityTraits int

const (
	PersonalityTraitUnspecified PersonalityTraits = iota
	PersonalityTraitIntrovert
	PersonalityTraitExtrovert
	PersonalityTraitAdventurous
	PersonalityTraitHomebody
	PersonalityTraitOptimist
	PersonalityTraitAmbitious
	PersonalityTraitCreative
	PersonalityTraitEmpathic
	PersonalityTraitAnalytical
	PersonalityTraitSarcasm
)

func PBToPersonalityTraits(pt desc.PersonalityTraits) PersonalityTraits {
	switch pt {
	case desc.PersonalityTraits_TRAIT_UNSPECIFIED:
		return PersonalityTraitUnspecified
	case desc.PersonalityTraits_TRAIT_INTROVERT:
		return PersonalityTraitIntrovert
	case desc.PersonalityTraits_TRAIT_EXTROVERT:
		return PersonalityTraitExtrovert
	case desc.PersonalityTraits_TRAIT_ADVENTUROUS:
		return PersonalityTraitAdventurous
	case desc.PersonalityTraits_TRAIT_HOMEBODY:
		return PersonalityTraitHomebody
	case desc.PersonalityTraits_TRAIT_OPTIMIST:
		return PersonalityTraitOptimist
	case desc.PersonalityTraits_TRAIT_AMBITIOUS:
		return PersonalityTraitAmbitious
	case desc.PersonalityTraits_TRAIT_CREATIVE:
		return PersonalityTraitCreative
	case desc.PersonalityTraits_TRAIT_EMPATHIC:
		return PersonalityTraitEmpathic
	case desc.PersonalityTraits_TRAIT_ANALYTICAL:
		return PersonalityTraitAnalytical
	case desc.PersonalityTraits_TRAIT_SARCASM:
		return PersonalityTraitSarcasm
	default:
		return PersonalityTraitUnspecified
	}
}

func PersonalityTraitsToPB(pt PersonalityTraits) desc.PersonalityTraits {
	switch pt {
	case PersonalityTraitUnspecified:
		return desc.PersonalityTraits_TRAIT_UNSPECIFIED
	case PersonalityTraitIntrovert:
		return desc.PersonalityTraits_TRAIT_INTROVERT
	case PersonalityTraitExtrovert:
		return desc.PersonalityTraits_TRAIT_EXTROVERT
	case PersonalityTraitAdventurous:
		return desc.PersonalityTraits_TRAIT_ADVENTUROUS
	case PersonalityTraitHomebody:
		return desc.PersonalityTraits_TRAIT_HOMEBODY
	case PersonalityTraitOptimist:
		return desc.PersonalityTraits_TRAIT_OPTIMIST
	case PersonalityTraitAmbitious:
		return desc.PersonalityTraits_TRAIT_AMBITIOUS
	case PersonalityTraitCreative:
		return desc.PersonalityTraits_TRAIT_CREATIVE
	case PersonalityTraitEmpathic:
		return desc.PersonalityTraits_TRAIT_EMPATHIC
	case PersonalityTraitAnalytical:
		return desc.PersonalityTraits_TRAIT_ANALYTICAL
	case PersonalityTraitSarcasm:
		return desc.PersonalityTraits_TRAIT_SARCASM
	default:
		return desc.PersonalityTraits_TRAIT_UNSPECIFIED
	}
}

func PersonalityTraitsSliceToPB(pt []PersonalityTraits) []desc.PersonalityTraits {
	pb := make([]desc.PersonalityTraits, len(pt))
	for i, v := range pt {
		pb[i] = PersonalityTraitsToPB(v)
	}
	return pb
}

func PBToPersonalityTraitsSlice(pb []desc.PersonalityTraits) []PersonalityTraits {
	pt := make([]PersonalityTraits, len(pb))
	for i, v := range pb {
		pt[i] = PBToPersonalityTraits(v)
	}
	return pt
}
