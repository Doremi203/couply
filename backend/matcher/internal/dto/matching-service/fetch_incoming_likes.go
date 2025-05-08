package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type FetchIncomingLikesV1Request struct {
	Limit  uint64
	Offset uint64
}

func (x *FetchIncomingLikesV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchIncomingLikesV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchIncomingLikesV1Response struct {
	Likes []*matching.Like
}

func (x *FetchIncomingLikesV1Response) GetLikes() []*matching.Like {
	if x != nil {
		return x.Likes
	}
	return nil
}

func PBToFetchIncomingLikesRequest(req *desc.FetchIncomingLikesV1Request) *FetchIncomingLikesV1Request {
	return &FetchIncomingLikesV1Request{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
}

func FetchIncomingLikesResponseToPB(resp *FetchIncomingLikesV1Response) *desc.FetchIncomingLikesV1Response {
	pbLikes := make([]*desc.Like, len(resp.GetLikes()))
	for i, like := range resp.GetLikes() {
		pbLikes[i] = matching.LikeToPB(like)
	}

	return &desc.FetchIncomingLikesV1Response{
		Likes: pbLikes,
	}
}
