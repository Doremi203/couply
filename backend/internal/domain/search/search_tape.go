package search

import desc "github.com/Doremi203/Couply/backend/pkg/search-service/v1"

type SearchTape struct {
	OnlyVerified bool
	OnlyPremium  bool
}

func SearchTapeToPB(st *SearchTape) *desc.SearchTape {
	return &desc.SearchTape{
		OnlyVerified: st.OnlyVerified,
		OnlyPremium:  st.OnlyPremium,
	}
}

func PBToSearchTape(st *desc.SearchTape) *SearchTape {
	return &SearchTape{
		OnlyVerified: st.OnlyVerified,
		OnlyPremium:  st.OnlyPremium,
	}
}
