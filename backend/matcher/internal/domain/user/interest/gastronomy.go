package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"

type Gastronomy int

const (
	GastronomyUnspecified Gastronomy = iota
	GastronomyCooking
	GastronomyWineDegustation
	GastronomyBars
	GastronomyCoffee
	GastronomyTea
	GastronomyVegan
	GastronomyFoodCritic
	GastronomySugarLover
)

func PBToGastronomy(gastronomy desc.Gastronomy) Gastronomy {
	switch gastronomy {
	case desc.Gastronomy_GASTRONOMY_UNSPECIFIED:
		return GastronomyUnspecified
	case desc.Gastronomy_GASTRONOMY_COOKING:
		return GastronomyCooking
	case desc.Gastronomy_GASTRONOMY_WINE_DEGUSTATION:
		return GastronomyWineDegustation
	case desc.Gastronomy_GASTRONOMY_BARS:
		return GastronomyBars
	case desc.Gastronomy_GASTRONOMY_COFFEE:
		return GastronomyCoffee
	case desc.Gastronomy_GASTRONOMY_TEA:
		return GastronomyTea
	case desc.Gastronomy_GASTRONOMY_VEGAN:
		return GastronomyVegan
	case desc.Gastronomy_GASTRONOMY_FOOD_CRITIC:
		return GastronomyFoodCritic
	case desc.Gastronomy_GASTRONOMY_SUGAR_LOVER:
		return GastronomySugarLover
	default:
		return Gastronomy(0)
	}
}

func GastronomyToPB(gastronomy Gastronomy) desc.Gastronomy {
	switch gastronomy {
	case GastronomyUnspecified:
		return desc.Gastronomy_GASTRONOMY_UNSPECIFIED
	case GastronomyCooking:
		return desc.Gastronomy_GASTRONOMY_COOKING
	case GastronomyWineDegustation:
		return desc.Gastronomy_GASTRONOMY_WINE_DEGUSTATION
	case GastronomyBars:
		return desc.Gastronomy_GASTRONOMY_BARS
	case GastronomyCoffee:
		return desc.Gastronomy_GASTRONOMY_COFFEE
	case GastronomyTea:
		return desc.Gastronomy_GASTRONOMY_TEA
	case GastronomyVegan:
		return desc.Gastronomy_GASTRONOMY_VEGAN
	case GastronomyFoodCritic:
		return desc.Gastronomy_GASTRONOMY_FOOD_CRITIC
	case GastronomySugarLover:
		return desc.Gastronomy_GASTRONOMY_SUGAR_LOVER
	default:
		return desc.Gastronomy(0)
	}
}

func GastronomySliceToPB(gastronomies []Gastronomy) []desc.Gastronomy {
	gastronomiesPB := make([]desc.Gastronomy, 0, len(gastronomies))

	for _, gastronomy := range gastronomies {
		gastronomiesPB = append(gastronomiesPB, GastronomyToPB(gastronomy))
	}

	return gastronomiesPB
}

func PBToGastronomySlice(gastronomies []desc.Gastronomy) []Gastronomy {
	gastronomiesDomain := make([]Gastronomy, 0, len(gastronomies))

	for _, gastronomy := range gastronomies {
		gastronomiesDomain = append(gastronomiesDomain, PBToGastronomy(gastronomy))
	}

	return gastronomiesDomain
}
