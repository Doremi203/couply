package dto

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

type ReportUserV1Request struct {
	BlockID       uuid.UUID
	TargetUserID  string
	ReportReasons []blocker.ReportReason
	Message       string
	CreatedAt     time.Time
}

type ReportUserV1Response struct{}

func ReportUserRequestToBlock(req *ReportUserV1Request) (*blocker.UserBlock, error) {
	blockedID, err := uuid.Parse(req.TargetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}
	return &blocker.UserBlock{
		ID:        req.BlockID,
		BlockedID: blockedID,
		Message:   req.Message,
		Reasons:   req.ReportReasons,
		Status:    blocker.BlockStatusPending,
		CreatedAt: req.CreatedAt,
	}, nil
}

func PBToReportUserRequest(req *desc.ReportUserV1Request) (*ReportUserV1Request, error) {
	reportReasons := make([]blocker.ReportReason, len(req.GetReasons()))
	for i, reason := range req.GetReasons() {
		reportReasons[i] = blocker.ReportReason(reason)
	}

	blockID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "uuid.NewV7")
	}

	return &ReportUserV1Request{
		BlockID:       blockID,
		TargetUserID:  req.GetTargetUserId(),
		ReportReasons: reportReasons,
		Message:       req.GetMessage(),
		CreatedAt:     time.Now(),
	}, nil
}

func ReportUserResponseToPB(_ *ReportUserV1Response) *desc.ReportUserV1Response {
	return &desc.ReportUserV1Response{}
}
