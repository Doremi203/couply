package matching_service

import (
	"fmt"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

type LikeUserV1Request struct {
	TargetUserId uuid.UUID
	Message      string
}

func (x *LikeUserV1Request) GetTargetUserId() uuid.UUID {
	if x != nil {
		return x.TargetUserId
	}
	return uuid.Nil
}

func (x *LikeUserV1Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LikeUserV1Response struct {
	IsMatch bool
	Match   *matching.Match
}

func (x *LikeUserV1Response) GetIsMatch() bool {
	if x != nil {
		return x.IsMatch
	}
	return false
}

func (x *LikeUserV1Response) GetMatch() *matching.Match {
	if x != nil {
		return x.Match
	}
	return nil
}

func PBToLikeUserRequest(req *desc.LikeUserV1Request) (*LikeUserV1Request, error) {
	targetUserID, err := uuid.Parse(req.GetTargetUserId())
	if err != nil {
		return nil, fmt.Errorf("invalid target user id: %s", req.GetTargetUserId())
	}

	return &LikeUserV1Request{
		TargetUserId: targetUserID,
		Message:      req.GetMessage(),
	}, nil
}

func LikeUserResponseToPB(resp *LikeUserV1Response) *desc.LikeUserV1Response {
	return &desc.LikeUserV1Response{
		IsMatch: resp.GetIsMatch(),
		Match:   matching.MatchToPB(resp.GetMatch()),
	}
}
