package dto

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

type GetBlockInfoV1Request struct{}

type GetBlockInfoV1Response struct {
	BlockID       uuid.UUID
	BlockedUserID uuid.UUID
	Message       string
	ReportReasons []blocker.ReportReason
	CreatedAt     time.Time
}

func (x *GetBlockInfoV1Response) GetBlockID() uuid.UUID {
	if x != nil {
		return x.BlockID
	}
	return uuid.Nil
}

func (x *GetBlockInfoV1Response) GetBlockedUserID() uuid.UUID {
	if x != nil {
		return x.BlockedUserID
	}
	return uuid.Nil
}

func (x *GetBlockInfoV1Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetBlockInfoV1Response) GetReportReasons() []blocker.ReportReason {
	if x != nil {
		return x.ReportReasons
	}
	return nil
}

func (x *GetBlockInfoV1Response) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func PBToGetBlockInfoRequest(_ *desc.GetBlockInfoV1Request) *GetBlockInfoV1Request {
	return &GetBlockInfoV1Request{}
}

func GetBlockInfoResponseToPB(resp *GetBlockInfoV1Response) *desc.GetBlockInfoV1Response {
	pbReasons := make([]desc.ReportReason, len(resp.GetReportReasons()))
	for i, reason := range resp.GetReportReasons() {
		pbReasons[i] = blocker.ReportReasonToPB(reason)
	}
	return &desc.GetBlockInfoV1Response{
		BlockId:       resp.GetBlockID().String(),
		BlockedUserId: resp.GetBlockedUserID().String(),
		Message:       resp.GetMessage(),
		Reasons:       pbReasons,
		CreatedAt:     timestamppb.New(resp.GetCreatedAt()),
	}
}
