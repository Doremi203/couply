package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type DislikeUserV1Request struct {
	TargetUserID uuid.UUID
}

type DislikeUserV1Response struct{}

func PBToDislikeUserRequest(req *desc.DislikeUserV1Request) (*DislikeUserV1Request, error) {
	targetUserID, err := uuid.Parse(req.GetTargetUserId())
	if err != nil {
		return nil, err
	}

	return &DislikeUserV1Request{
		TargetUserID: targetUserID,
	}, nil
}

func DislikeUserResponseToPB(_ *DislikeUserV1Response) *desc.DislikeUserV1Response {
	return &desc.DislikeUserV1Response{}
}
