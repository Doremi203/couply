package matching_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type UpdateMatchV1Request struct {
	MainUserID   int64
	ChosenUserID int64
	Approved     bool
}

type UpdateMatchV1Response struct {
	Match *matching.Match
}

func UpdateMatchRequestToPB(req *UpdateMatchV1Request) *desc.UpdateMatchV1Request {
	return &desc.UpdateMatchV1Request{
		MainUserId:   req.MainUserID,
		ChosenUserId: req.ChosenUserID,
		Approved:     req.Approved,
	}
}

func PBToUpdateMatchRequest(req *desc.UpdateMatchV1Request) *UpdateMatchV1Request {
	return &UpdateMatchV1Request{
		MainUserID:   req.GetMainUserId(),
		ChosenUserID: req.GetChosenUserId(),
		Approved:     req.GetApproved(),
	}
}

func UpdateMatchResponseToPB(resp *UpdateMatchV1Response) *desc.UpdateMatchV1Response {
	return &desc.UpdateMatchV1Response{
		Match: matching.MatchToPB(resp.Match),
	}
}

func PBToUpdateMatchResponse(resp *desc.UpdateMatchV1Response) *UpdateMatchV1Response {
	return &UpdateMatchV1Response{
		Match: matching.PBToMatch(resp.GetMatch()),
	}
}

func UpdateMatchRequestToMatch(req *UpdateMatchV1Request) *matching.Match {
	return &matching.Match{
		MainUserID:   req.MainUserID,
		ChosenUserID: req.ChosenUserID,
		Approved:     req.Approved,
	}
}
