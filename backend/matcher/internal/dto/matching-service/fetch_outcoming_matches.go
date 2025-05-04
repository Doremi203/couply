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

func (x *FetchOutgoingMatchesV1Request) GetMainUserID() int64 {
	if x != nil {
		return x.MainUserID
	}
	return 0
}

func (x *FetchOutgoingMatchesV1Request) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchOutgoingMatchesV1Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchOutgoingMatchesV1Response struct {
	Matches []*matching.Match
}

func (x *FetchOutgoingMatchesV1Response) GetMatches() []*matching.Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func FetchOutgoingMatchesRequestToPB(req *FetchOutgoingMatchesV1Request) *desc.FetchOutgoingMatchesV1Request {
	return &desc.FetchOutgoingMatchesV1Request{
		MainUserId: req.GetMainUserID(),
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
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
		Match: matching.MatchSliceToPB(resp.GetMatches()),
	}
}

func PBToFetchOutgoingMatchesResponse(resp *desc.FetchOutgoingMatchesV1Response) *FetchOutgoingMatchesV1Response {
	return &FetchOutgoingMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.GetMatch()),
	}
}
