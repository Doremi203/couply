package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type FetchMatchesUserIDsV1Request struct {
	Limit  uint64
	Offset uint64
}

func (x *FetchMatchesUserIDsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchMatchesUserIDsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchMatchesUserIDsV1Response struct {
	UserIDs []*uuid.UUID
}

func (x *FetchMatchesUserIDsV1Response) GetUserIDs() []*uuid.UUID {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func PBToFetchMatchesUserIDsRequest(req *desc.FetchMatchesUserIDsV1Request) *FetchMatchesUserIDsV1Request {
	return &FetchMatchesUserIDsV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchMatchesUserIDsResponseToPB(resp *FetchMatchesUserIDsV1Response) *desc.FetchMatchesUserIDsV1Response {
	pbUserIDs := make([]string, len(resp.GetUserIDs()))
	for i, id := range resp.GetUserIDs() {
		pbUserIDs[i] = id.String()
	}

	return &desc.FetchMatchesUserIDsV1Response{
		UserIds: pbUserIDs,
	}
}
