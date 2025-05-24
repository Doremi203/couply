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

type LikeUserV1Response struct {
	IsMatch bool
	Match   *matching.Match
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
		IsMatch: resp.IsMatch,
		Match:   matching.MatchToPB(resp.Match),
	}
}
