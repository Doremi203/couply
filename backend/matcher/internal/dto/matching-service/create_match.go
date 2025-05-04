package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type CreateMatchV1Request struct {
	MainUserID   int64
	ChosenUserID int64
}

func (x *CreateMatchV1Request) GetMainUserID() int64 {
	if x != nil {
		return x.MainUserID
	}
	return 0
}

func (x *CreateMatchV1Request) GetChosenUserID() int64 {
	if x != nil {
		return x.ChosenUserID
	}
	return 0
}

type CreateMatchV1Response struct {
	Match *matching.Match
}

func (x *CreateMatchV1Response) GetMatch() *matching.Match {
	if x != nil {
		return x.Match
	}
	return nil
}

func CreateMatchRequestToPB(req *CreateMatchV1Request) *desc.CreateMatchV1Request {
	return &desc.CreateMatchV1Request{
		MainUserId:   req.GetMainUserID(),
		ChosenUserId: req.GetChosenUserID(),
	}
}

func PBToCreateMatchRequest(req *desc.CreateMatchV1Request) *CreateMatchV1Request {
	return &CreateMatchV1Request{
		MainUserID:   req.GetMainUserId(),
		ChosenUserID: req.GetChosenUserId(),
	}
}

func CreateMatchResponseToPB(resp *CreateMatchV1Response) *desc.CreateMatchV1Response {
	return &desc.CreateMatchV1Response{
		Match: matching.MatchToPB(resp.GetMatch()),
	}
}

func PBToCreateMatchResponse(resp *desc.CreateMatchV1Response) *CreateMatchV1Response {
	return &CreateMatchV1Response{
		Match: matching.PBToMatch(resp.GetMatch()),
	}
}

func CreateMatchRequestToMatch(req *CreateMatchV1Request) *matching.Match {
	return &matching.Match{
		MainUserID:   req.GetMainUserID(),
		ChosenUserID: req.GetChosenUserID(),
	}
}
