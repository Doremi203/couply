package matching_service

import (
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
	desc "github.com/Doremi203/Couply/backend/pkg/matching-service/v1"
)

type CreateMatchV1Request struct {
	MainUserID   int64
	ChosenUserID int64
}

type CreateMatchV1Response struct {
	Match *matching.Match
}

func CreateMatchRequestToPB(req *CreateMatchV1Request) *desc.CreateMatchV1Request {
	return &desc.CreateMatchV1Request{
		MainUserId:   req.MainUserID,
		ChosenUserId: req.ChosenUserID,
	}
}

func PBToCreateMatchRequest(req *desc.CreateMatchV1Request) *CreateMatchV1Request {
	return &CreateMatchV1Request{
		MainUserID:   req.MainUserId,
		ChosenUserID: req.ChosenUserId,
	}
}

func CreateMatchResponseToPB(resp *CreateMatchV1Response) *desc.CreateMatchV1Response {
	return &desc.CreateMatchV1Response{
		Match: matching.MatchToPB(resp.Match),
	}
}

func PBToCreateMatchResponse(resp *desc.CreateMatchV1Response) *CreateMatchV1Response {
	return &CreateMatchV1Response{
		Match: matching.PBToMatch(resp.Match),
	}
}

func CreateMatchRequestToMatch(req *CreateMatchV1Request) *matching.Match {
	return &matching.Match{
		MainUserID:   req.MainUserID,
		ChosenUserID: req.ChosenUserID,
	}
}
