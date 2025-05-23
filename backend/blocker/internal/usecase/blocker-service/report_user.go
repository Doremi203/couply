package blocker_service

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserV1Request) (*dto.ReportUserV1Response, error) {
	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.TargetUserID)
	if err != nil {
		return nil, err
	}

	blockID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	err = c.bot.SendReportMessage(reportedUser, in.ReportReasons, in.Message, blockID)
	if err != nil {
		return nil, err
	}

	targetUserID, err := uuid.Parse(in.TargetUserID)
	if err != nil {
		return nil, err
	}

	block := &blocker.UserBlock{
		ID:        blockID,
		BlockedID: targetUserID,
		Message:   in.Message,
		Reasons:   in.ReportReasons,
		Status:    blocker.BlockStatusPending,
		CreatedAt: time.Now(),
	}

	err = c.blockerStorageFacade.ReportUserTx(ctx, block)
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserV1Response{}, nil
}
