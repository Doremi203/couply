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
	BlockStatus   blocker.BlockStatus
	CreatedAt     time.Time
}

func PBToGetBlockInfoRequest(_ *desc.GetBlockInfoV1Request) *GetBlockInfoV1Request {
	return &GetBlockInfoV1Request{}
}

func GetBlockInfoResponseToPB(resp *GetBlockInfoV1Response) *desc.GetBlockInfoV1Response {
	pbReasons := make([]desc.ReportReason, len(resp.ReportReasons))
	for i, reason := range resp.ReportReasons {
		pbReasons[i] = blocker.ReportReasonToPB(reason)
	}
	return &desc.GetBlockInfoV1Response{
		BlockId:       resp.BlockID.String(),
		BlockedUserId: resp.BlockedUserID.String(),
		Message:       resp.Message,
		Reasons:       pbReasons,
		Status:        blocker.BlockStatusToPB(resp.BlockStatus),
		CreatedAt:     timestamppb.New(resp.CreatedAt),
	}
}
