package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type GetFilterV1Request struct {
	UserID int64
}

type GetFilterV1Response struct {
	Filter *search.Filter
}

func GetFilterRequestToPB(req *GetFilterV1Request) *desc.GetFilterV1Request {
	return &desc.GetFilterV1Request{
		UserId: req.UserID,
	}
}

func PBToGetFilterRequest(req *desc.GetFilterV1Request) *GetFilterV1Request {
	return &GetFilterV1Request{
		UserID: req.GetUserId(),
	}
}

func GetFilterResponseToPB(resp *GetFilterV1Response) *desc.GetFilterV1Response {
	return &desc.GetFilterV1Response{
		Filter: search.FilterToPB(resp.Filter),
	}
}

func PBToGetFilterResponse(resp *desc.GetFilterV1Response) *GetFilterV1Response {
	return &GetFilterV1Response{
		Filter: search.PBToFilter(resp.GetFilter()),
	}
}
