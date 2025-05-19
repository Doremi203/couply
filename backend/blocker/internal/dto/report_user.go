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

func (x *ReportUserV1Request) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *ReportUserV1Request) GetReportReasons() []blocker.ReportReason {
	if x != nil {
		return x.ReportReasons
	}
	return nil
}

func (x *ReportUserV1Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReportUserV1Response struct{}

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
