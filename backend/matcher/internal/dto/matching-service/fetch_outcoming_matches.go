package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchOutgoingMatchesV1Request struct {
	MainUserID int64
	Limit      int32
	Offset     int32
}

type FetchOutgoingMatchesV1Response struct {
	Matches []*matching.Match
}

func FetchOutgoingMatchesRequestToPB(req *FetchOutgoingMatchesV1Request) *desc.FetchOutgoingMatchesV1Request {
	return &desc.FetchOutgoingMatchesV1Request{
		MainUserId: req.MainUserID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}
}

func PBToFetchOutgoingMatchesRequest(req *desc.FetchOutgoingMatchesV1Request) *FetchOutgoingMatchesV1Request {
	return &FetchOutgoingMatchesV1Request{
		MainUserID: req.GetMainUserId(),
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	}
}

func FetchOutgoingMatchesResponseToPB(resp *FetchOutgoingMatchesV1Response) *desc.FetchOutgoingMatchesV1Response {
	return &desc.FetchOutgoingMatchesV1Response{
		Match: matching.MatchSliceToPB(resp.Matches),
	}
}

func PBToFetchOutgoingMatchesResponseSlice(resp *desc.FetchOutgoingMatchesV1Response) *FetchOutgoingMatchesV1Response {
	return &FetchOutgoingMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.GetMatch()),
	}
}
