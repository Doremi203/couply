package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type GetFilterV1Request struct {
}

type GetFilterV1Response struct {
	Filter *search.Filter
}

func (x *GetFilterV1Response) GetFilter() *search.Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func PBToGetFilterRequest(_ *desc.GetFilterV1Request) *GetFilterV1Request {
	return &GetFilterV1Request{}
}

func GetFilterResponseToPB(resp *GetFilterV1Response) *desc.GetFilterV1Response {
	return &desc.GetFilterV1Response{
		Filter: search.FilterToPB(resp.GetFilter()),
	}
}
