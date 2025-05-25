package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type FetchMatchesUserIDsV1Request struct {
	Limit  uint64
	Offset uint64
}

type FetchMatchesUserIDsV1Response struct {
	UserIDs []*uuid.UUID
}

func PBToFetchMatchesUserIDsRequest(req *desc.FetchMatchesUserIDsV1Request) *FetchMatchesUserIDsV1Request {
	return &FetchMatchesUserIDsV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchMatchesUserIDsResponseToPB(resp *FetchMatchesUserIDsV1Response) *desc.FetchMatchesUserIDsV1Response {
	pbUserIDs := make([]string, len(resp.UserIDs))
	for i, id := range resp.UserIDs {
		pbUserIDs[i] = id.String()
	}

	return &desc.FetchMatchesUserIDsV1Response{
		UserIds: pbUserIDs,
	}
}
