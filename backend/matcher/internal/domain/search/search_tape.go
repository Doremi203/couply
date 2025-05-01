package search

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"

type SearchPreferences struct {
	OnlyVerified bool
	OnlyPremium  bool
}

func SearchPreferencesToPB(st *SearchPreferences) *desc.SearchPreferences {
	return &desc.SearchPreferences{
		OnlyVerified: st.OnlyVerified,
		OnlyPremium:  st.OnlyPremium,
	}
}

func PBToSearchPreferences(st *desc.SearchPreferences) *SearchPreferences {
	return &SearchPreferences{
		OnlyVerified: st.OnlyVerified,
		OnlyPremium:  st.OnlyPremium,
	}
}
