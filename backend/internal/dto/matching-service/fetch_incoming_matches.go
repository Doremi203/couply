package matching_service

import (
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
	desc "github.com/Doremi203/Couply/backend/pkg/matching-service/v1"
)

type FetchIncomingMatchesV1Request struct {
	ChosenUserID int64
	Limit        int32
	Offset       int32
}

type FetchIncomingMatchesV1Response struct {
	Matches []*matching.Match
}

func FetchIncomingMatchesRequestToPB(req *FetchIncomingMatchesV1Request) *desc.FetchIncomingMatchesV1Request {
	return &desc.FetchIncomingMatchesV1Request{
		ChosenUserId: req.ChosenUserID,
		Limit:        req.Limit,
		Offset:       req.Offset,
	}
}

func PBToFetchIncomingMatchesRequest(req *desc.FetchIncomingMatchesV1Request) *FetchIncomingMatchesV1Request {
	return &FetchIncomingMatchesV1Request{
		ChosenUserID: req.ChosenUserId,
		Limit:        req.Limit,
		Offset:       req.Offset,
	}
}

func FetchIncomingMatchesResponseToPB(resp *FetchIncomingMatchesV1Response) *desc.FetchIncomingMatchesV1Response {
	return &desc.FetchIncomingMatchesV1Response{
		Match: matching.MatchSliceToPB(resp.Matches),
	}
}

func PBToFetchIncomingMatchesResponseSlice(resp *desc.FetchIncomingMatchesV1Response) *FetchIncomingMatchesV1Response {
	return &FetchIncomingMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.Match),
	}
}
