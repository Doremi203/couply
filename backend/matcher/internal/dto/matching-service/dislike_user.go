package matching_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
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
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	return &DislikeUserV1Request{
		TargetUserID: targetUserID,
	}, nil
}

func DislikeUserResponseToPB(_ *DislikeUserV1Response) *desc.DislikeUserV1Response {
	return &desc.DislikeUserV1Response{}
}
