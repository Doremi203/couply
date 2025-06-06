package matching_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type DeleteMatchV1Request struct {
	TargetUserID uuid.UUID
}

type DeleteMatchV1Response struct{}

func PBToDeleteMatchRequest(req *desc.DeleteMatchV1Request) (*DeleteMatchV1Request, error) {
	targetUserID, err := uuid.Parse(req.GetTargetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	return &DeleteMatchV1Request{
		TargetUserID: targetUserID,
	}, nil
}

func DeleteMatchResponseToPB(_ *DeleteMatchV1Response) *desc.DeleteMatchV1Response {
	return &desc.DeleteMatchV1Response{}
}
