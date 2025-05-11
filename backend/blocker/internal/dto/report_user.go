package dto

import (
	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
)

type ReportUserRequest struct {
	UserID        string
	ReportReasons []blocker.ReportReason
	Message       string
}

func (x *ReportUserRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ReportUserRequest) GetReportReasons() []blocker.ReportReason {
	if x != nil {
		return x.ReportReasons
	}
	return nil
}

func (x *ReportUserRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReportUserResponse struct{}

func PBToReportUserRequest(req *desc.ReportUserV1Request) *ReportUserRequest {
	reportReasons := make([]blocker.ReportReason, len(req.GetReasons()))
	for i, reason := range req.GetReasons() {
		reportReasons[i] = blocker.ReportReason(reason)
	}

	return &ReportUserRequest{
		UserID:        req.GetUserId(),
		ReportReasons: reportReasons,
		Message:       req.GetMessage(),
	}
}

func ReportUserResponseToPB(resp *ReportUserResponse) *desc.ReportUserV1Response {
	return &desc.ReportUserV1Response{}
}
