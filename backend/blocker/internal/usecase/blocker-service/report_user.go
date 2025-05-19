package blocker_service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserV1Request) (*dto.ReportUserV1Response, error) {
	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.GetTargetUserID())
	if err != nil {
		return nil, err
	}

	err = c.bot.SendReportMessage(reportedUser, in.GetReportReasons(), in.GetMessage())
	if err != nil {
		return nil, err
	}

	blockID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	targetUserID, err := uuid.Parse(in.GetTargetUserID())
	if err != nil {
		return nil, err
	}

	err = c.blockerStorageFacade.ReportUserTx(ctx, blockID, targetUserID, in.GetMessage(), time.Now(), in.GetReportReasons())
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserV1Response{}, nil
}
