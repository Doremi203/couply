package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchOutgoingLikesV1Request struct {
	Limit  uint64
	Offset uint64
}

type FetchOutgoingLikesV1Response struct {
	Likes []*matching.Like
}

func PBToFetchOutgoingLikesRequest(req *desc.FetchOutgoingLikesV1Request) *FetchOutgoingLikesV1Request {
	return &FetchOutgoingLikesV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchOutgoingLikesResponseToPB(resp *FetchOutgoingLikesV1Response) *desc.FetchOutgoingLikesV1Response {
	pbLikes := make([]*desc.Like, len(resp.Likes))
	for i, like := range resp.Likes {
		pbLikes[i] = matching.LikeToPB(like)
	}

	return &desc.FetchOutgoingLikesV1Response{
		Likes: pbLikes,
	}
}
