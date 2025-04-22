package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchMatchesV1Request struct {
	MainUserID int64
	Limit      int32
	Offset     int32
}

type FetchMatchesV1Response struct {
	Matches []*matching.Match
}

func FetchMatchesRequestToPB(req *FetchMatchesV1Request) *desc.FetchMatchesV1Request {
	return &desc.FetchMatchesV1Request{
		MainUserId: req.MainUserID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}
}

func PBToFetchMatchesRequest(req *desc.FetchMatchesV1Request) *FetchMatchesV1Request {
	return &FetchMatchesV1Request{
		MainUserID: req.MainUserId,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}
}

func FetchMatchesResponseToPB(resp *FetchMatchesV1Response) *desc.FetchMatchesV1Response {
	return &desc.FetchMatchesV1Response{
		Match: matching.MatchSliceToPB(resp.Matches),
	}
}

func PBToFetchMatchesResponseSlice(resp *desc.FetchMatchesV1Response) *FetchMatchesV1Response {
	return &FetchMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.Match),
	}
}
