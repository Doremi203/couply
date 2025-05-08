package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type FetchMatchesV1Request struct {
	Limit  uint64
	Offset uint64
}

func (x *FetchMatchesV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchMatchesV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchMatchesV1Response struct {
	UserIDs []*uuid.UUID
}

func (x *FetchMatchesV1Response) GetUserIDs() []*uuid.UUID {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func PBToFetchMatchesRequest(req *desc.FetchMatchesUserIDsV1Request) *FetchMatchesV1Request {
	return &FetchMatchesV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchMatchesResponseToPB(resp *FetchMatchesV1Response) *desc.FetchMatchesUserIDsV1Response {
	pbUserIDs := make([]string, len(resp.GetUserIDs()))
	for i, id := range resp.GetUserIDs() {
		pbUserIDs[i] = id.String()
	}

	return &desc.FetchMatchesUserIDsV1Response{
		UserIds: pbUserIDs,
	}
}
