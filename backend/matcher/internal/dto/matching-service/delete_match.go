package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type DeleteMatchV1Request struct {
	TargetUserID uuid.UUID
}

func (x *DeleteMatchV1Request) GetTargetUserID() uuid.UUID {
	if x != nil {
		return x.TargetUserID
	}
	return uuid.Nil
}

type DeleteMatchV1Response struct{}

func PBToDeleteMatchRequest(req *desc.DeleteMatchV1Request) (*DeleteMatchV1Request, error) {
	targetUserID, err := uuid.Parse(req.GetTargetUserId())
	if err != nil {
		return nil, err
	}

	return &DeleteMatchV1Request{
		TargetUserID: targetUserID,
	}, nil
}

func DeleteMatchResponseToPB(_ *DeleteMatchV1Response) *desc.DeleteMatchV1Response {
	return &desc.DeleteMatchV1Response{}
}
