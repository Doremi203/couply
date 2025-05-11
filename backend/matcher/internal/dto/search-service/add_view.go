package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/google/uuid"
)

type AddViewV1Request struct {
	ViewedID uuid.UUID
}

func (x *AddViewV1Request) GetViewedID() uuid.UUID {
	if x != nil {
		return x.ViewedID
	}
	return uuid.Nil
}

type AddViewV1Response struct{}

func PBToAddViewRequest(req *desc.AddViewV1Request) (*AddViewV1Request, error) {
	viewedID, err := uuid.Parse(req.GetViewedId())
	if err != nil {
		return nil, err
	}

	return &AddViewV1Request{
		ViewedID: viewedID,
	}, nil
}

func AddViewResponseToPB(_ *AddViewV1Response) *desc.AddViewV1Response {
	return &desc.AddViewV1Response{}
}
