package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchOutgoingLikesV1Request struct {
	Limit  uint64
	Offset uint64
}

func (x *FetchOutgoingLikesV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchOutgoingLikesV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchOutgoingLikesV1Response struct {
	Likes []*matching.Like
}

func (x *FetchOutgoingLikesV1Response) GetLikes() []*matching.Like {
	if x != nil {
		return x.Likes
	}
	return nil
}

func PBToFetchOutgoingLikesRequest(req *desc.FetchOutgoingLikesV1Request) *FetchOutgoingLikesV1Request {
	return &FetchOutgoingLikesV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchOutgoingLikesResponseToPB(resp *FetchOutgoingLikesV1Response) *desc.FetchOutgoingLikesV1Response {
	pbLikes := make([]*desc.Like, len(resp.GetLikes()))
	for i, like := range resp.GetLikes() {
		pbLikes[i] = matching.LikeToPB(like)
	}

	return &desc.FetchOutgoingLikesV1Response{
		Likes: pbLikes,
	}
}
