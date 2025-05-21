package blocker_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserV1Request) (*dto.ReportUserV1Response, error) {
	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.GetTargetUserID())
	if err != nil {
		return nil, err
	}

	blockID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	err = c.bot.SendReportMessage(reportedUser, in.GetReportReasons(), in.GetMessage(), blockID)
	if err != nil {
		return nil, err
	}

	targetUserID, err := uuid.Parse(in.GetTargetUserID())
	if err != nil {
		return nil, err
	}

	block := &blocker.UserBlock{
		ID:        blockID,
		BlockedID: targetUserID,
		Message:   in.GetMessage(),
		Reasons:   in.GetReportReasons(),
		Status:    blocker.BlockStatusPending,
		CreatedAt: time.Now(),
	}

	err = c.blockerStorageFacade.ReportUserTx(ctx, block)
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserV1Response{}, nil
}
