package dto

import (
	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
)

type ReportUserV1Request struct {
	TargetUserID  string
	ReportReasons []blocker.ReportReason
	Message       string
}

type ReportUserV1Response struct{}

func PBToReportUserRequest(req *desc.ReportUserV1Request) *ReportUserV1Request {
	reportReasons := make([]blocker.ReportReason, len(req.Reasons))
	for i, reason := range req.Reasons {
		reportReasons[i] = blocker.ReportReason(reason)
	}

	return &ReportUserV1Request{
		TargetUserID:  req.TargetUserId,
		ReportReasons: reportReasons,
		Message:       req.Message,
	}
}

func ReportUserResponseToPB(_ *ReportUserV1Response) *desc.ReportUserV1Response {
	return &desc.ReportUserV1Response{}
}
