package interest

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/common/v1"

type FoodDrink int

const (
	FoodDrinkUnspecified FoodDrink = iota
	FoodDrinkCoffee
	FoodDrinkWine
	FoodDrinkCocktails
	FoodDrinkVegan
	FoodDrinkBaking
	FoodDrinkFineDining
	FoodDrinkStreetFood
	FoodDrinkTea
	FoodDrinkBarbecue
	FoodDrinkCraftBeer
)

func PBToFoodDrink(fd desc.FoodDrink) FoodDrink {
	switch fd {
	case desc.FoodDrink_FOODDRINK_UNSPECIFIED:
		return FoodDrinkUnspecified
	case desc.FoodDrink_FOODDRINK_COFFEE:
		return FoodDrinkCoffee
	case desc.FoodDrink_FOODDRINK_WINE:
		return FoodDrinkWine
	case desc.FoodDrink_FOODDRINK_COCKTAILS:
		return FoodDrinkCocktails
	case desc.FoodDrink_FOODDRINK_VEGAN:
		return FoodDrinkVegan
	case desc.FoodDrink_FOODDRINK_BAKING:
		return FoodDrinkBaking
	case desc.FoodDrink_FOODDRINK_FINE_DINING:
		return FoodDrinkFineDining
	case desc.FoodDrink_FOODDRINK_STREET_FOOD:
		return FoodDrinkStreetFood
	case desc.FoodDrink_FOODDRINK_TEA:
		return FoodDrinkTea
	case desc.FoodDrink_FOODDRINK_BARBECUE:
		return FoodDrinkBarbecue
	case desc.FoodDrink_FOODDRINK_CRAFT_BEER:
		return FoodDrinkCraftBeer
	default:
		return FoodDrinkUnspecified
	}
}

func FoodDrinkToPB(fd FoodDrink) desc.FoodDrink {
	switch fd {
	case FoodDrinkUnspecified:
		return desc.FoodDrink_FOODDRINK_UNSPECIFIED
	case FoodDrinkCoffee:
		return desc.FoodDrink_FOODDRINK_COFFEE
	case FoodDrinkWine:
		return desc.FoodDrink_FOODDRINK_WINE
	case FoodDrinkCocktails:
		return desc.FoodDrink_FOODDRINK_COCKTAILS
	case FoodDrinkVegan:
		return desc.FoodDrink_FOODDRINK_VEGAN
	case FoodDrinkBaking:
		return desc.FoodDrink_FOODDRINK_BAKING
	case FoodDrinkFineDining:
		return desc.FoodDrink_FOODDRINK_FINE_DINING
	case FoodDrinkStreetFood:
		return desc.FoodDrink_FOODDRINK_STREET_FOOD
	case FoodDrinkTea:
		return desc.FoodDrink_FOODDRINK_TEA
	case FoodDrinkBarbecue:
		return desc.FoodDrink_FOODDRINK_BARBECUE
	case FoodDrinkCraftBeer:
		return desc.FoodDrink_FOODDRINK_CRAFT_BEER
	default:
		return desc.FoodDrink_FOODDRINK_UNSPECIFIED
	}
}

func FoodDrinkSliceToPB(fd []FoodDrink) []desc.FoodDrink {
	pb := make([]desc.FoodDrink, len(fd))
	for i, v := range fd {
		pb[i] = FoodDrinkToPB(v)
	}
	return pb
}

func PBToFoodDrinkSlice(pb []desc.FoodDrink) []FoodDrink {
	fd := make([]FoodDrink, len(pb))
	for i, v := range pb {
		fd[i] = PBToFoodDrink(v)
	}
	return fd
}
