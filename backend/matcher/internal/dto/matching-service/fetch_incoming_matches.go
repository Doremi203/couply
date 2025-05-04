package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchIncomingMatchesV1Request struct {
	ChosenUserID int64
	Limit        int32
	Offset       int32
}

func (x *FetchIncomingMatchesV1Request) GetChosenUserID() int64 {
	if x != nil {
		return x.ChosenUserID
	}
	return 0
}

func (x *FetchIncomingMatchesV1Request) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchIncomingMatchesV1Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchIncomingMatchesV1Response struct {
	Matches []*matching.Match
}

func (x *FetchIncomingMatchesV1Response) GetMatches() []*matching.Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func FetchIncomingMatchesRequestToPB(req *FetchIncomingMatchesV1Request) *desc.FetchIncomingMatchesV1Request {
	return &desc.FetchIncomingMatchesV1Request{
		ChosenUserId: req.GetChosenUserID(),
		Limit:        req.GetLimit(),
		Offset:       req.GetOffset(),
	}
}

func PBToFetchIncomingMatchesRequest(req *desc.FetchIncomingMatchesV1Request) *FetchIncomingMatchesV1Request {
	return &FetchIncomingMatchesV1Request{
		ChosenUserID: req.GetChosenUserId(),
		Limit:        req.GetLimit(),
		Offset:       req.GetOffset(),
	}
}

func FetchIncomingMatchesResponseToPB(resp *FetchIncomingMatchesV1Response) *desc.FetchIncomingMatchesV1Response {
	return &desc.FetchIncomingMatchesV1Response{
		Match: matching.MatchSliceToPB(resp.GetMatches()),
	}
}

func PBToFetchIncomingMatchesResponse(resp *desc.FetchIncomingMatchesV1Response) *FetchIncomingMatchesV1Response {
	return &FetchIncomingMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.GetMatch()),
	}
}
