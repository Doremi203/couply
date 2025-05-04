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

func (x *FetchMatchesV1Request) GetMainUserID() int64 {
	if x != nil {
		return x.MainUserID
	}
	return 0
}

func (x *FetchMatchesV1Request) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchMatchesV1Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchMatchesV1Response struct {
	Matches []*matching.Match
}

func (x *FetchMatchesV1Response) GetMatches() []*matching.Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func FetchMatchesRequestToPB(req *FetchMatchesV1Request) *desc.FetchMatchesV1Request {
	return &desc.FetchMatchesV1Request{
		MainUserId: req.GetMainUserID(),
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	}
}

func PBToFetchMatchesRequest(req *desc.FetchMatchesV1Request) *FetchMatchesV1Request {
	return &FetchMatchesV1Request{
		MainUserID: req.GetMainUserId(),
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	}
}

func FetchMatchesResponseToPB(resp *FetchMatchesV1Response) *desc.FetchMatchesV1Response {
	return &desc.FetchMatchesV1Response{
		Match: matching.MatchSliceToPB(resp.GetMatches()),
	}
}

func PBToFetchMatchesResponse(resp *desc.FetchMatchesV1Response) *FetchMatchesV1Response {
	return &FetchMatchesV1Response{
		Matches: matching.PBToMatchSlice(resp.GetMatch()),
	}
}
