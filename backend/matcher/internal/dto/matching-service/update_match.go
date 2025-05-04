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

func (x *UpdateMatchV1Request) GetMainUserID() int64 {
	if x != nil {
		return x.MainUserID
	}
	return 0
}

func (x *UpdateMatchV1Request) GetChosenUserID() int64 {
	if x != nil {
		return x.ChosenUserID
	}
	return 0
}

func (x *UpdateMatchV1Request) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

type UpdateMatchV1Response struct {
	Match *matching.Match
}

func (x *UpdateMatchV1Response) GetMatch() *matching.Match {
	if x != nil {
		return x.Match
	}
	return nil
}

func UpdateMatchRequestToPB(req *UpdateMatchV1Request) *desc.UpdateMatchV1Request {
	return &desc.UpdateMatchV1Request{
		MainUserId:   req.GetMainUserID(),
		ChosenUserId: req.GetChosenUserID(),
		Approved:     req.GetApproved(),
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
		Match: matching.MatchToPB(resp.GetMatch()),
	}
}

func PBToUpdateMatchResponse(resp *desc.UpdateMatchV1Response) *UpdateMatchV1Response {
	return &UpdateMatchV1Response{
		Match: matching.PBToMatch(resp.GetMatch()),
	}
}

func UpdateMatchRequestToMatch(req *UpdateMatchV1Request) *matching.Match {
	return &matching.Match{
		MainUserID:   req.GetMainUserID(),
		ChosenUserID: req.GetChosenUserID(),
		Approved:     req.GetApproved(),
	}
}
