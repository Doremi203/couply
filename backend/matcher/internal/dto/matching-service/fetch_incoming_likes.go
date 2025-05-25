package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchIncomingLikesV1Request struct {
	Limit  uint64
	Offset uint64
}

type FetchIncomingLikesV1Response struct {
	Likes []*matching.Like
}

func PBToFetchIncomingLikesRequest(req *desc.FetchIncomingLikesV1Request) *FetchIncomingLikesV1Request {
	return &FetchIncomingLikesV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchIncomingLikesResponseToPB(resp *FetchIncomingLikesV1Response) *desc.FetchIncomingLikesV1Response {
	pbLikes := make([]*desc.Like, len(resp.Likes))
	for i, like := range resp.Likes {
		pbLikes[i] = matching.LikeToPB(like)
	}

	return &desc.FetchIncomingLikesV1Response{
		Likes: pbLikes,
	}
}
