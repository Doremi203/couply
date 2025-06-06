package search_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/google/uuid"
)

type AddViewV1Request struct {
	ViewedID uuid.UUID
}

type AddViewV1Response struct{}

func PBToAddViewRequest(req *desc.AddViewV1Request) (*AddViewV1Request, error) {
	viewedID, err := uuid.Parse(req.GetViewedId())
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	return &AddViewV1Request{
		ViewedID: viewedID,
	}, nil
}

func AddViewResponseToPB(_ *AddViewV1Response) *desc.AddViewV1Response {
	return &desc.AddViewV1Response{}
}
