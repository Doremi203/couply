package dto

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

type ReportUserV1Request struct {
	TargetUserID  string
	ReportReasons []blocker.ReportReason
	Message       string
}

type ReportUserV1Response struct{}

func ReportUserRequestToBlock(req *ReportUserV1Request) (*blocker.UserBlock, error) {
	blockID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "uuid.NewV7")
	}

	blockedID, err := uuid.Parse(req.TargetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}
	return &blocker.UserBlock{
		ID:        blockID,
		BlockedID: blockedID,
		Message:   req.Message,
		Reasons:   req.ReportReasons,
		Status:    blocker.BlockStatusPending,
		CreatedAt: time.Now(),
	}, nil
}

func PBToReportUserRequest(req *desc.ReportUserV1Request) *ReportUserV1Request {
	reportReasons := make([]blocker.ReportReason, len(req.GetReasons()))
	for i, reason := range req.GetReasons() {
		reportReasons[i] = blocker.ReportReason(reason)
	}

	return &ReportUserV1Request{
		TargetUserID:  req.GetTargetUserId(),
		ReportReasons: reportReasons,
		Message:       req.GetMessage(),
	}
}

func ReportUserResponseToPB(_ *ReportUserV1Response) *desc.ReportUserV1Response {
	return &desc.ReportUserV1Response{}
}
