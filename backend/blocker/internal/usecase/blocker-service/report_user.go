package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserV1Request) (*dto.ReportUserV1Response, error) {
	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.TargetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "userServiceClient.GetUserByIDV1")
	}

	block, err := dto.ReportUserRequestToBlock(in)
	if err != nil {
		return nil, errors.Wrap(err, "dto.ReportUserRequestToBlock")
	}

	err = c.bot.SendReportMessage(reportedUser, block.Reasons, block.Message, block.ID)
	if err != nil {
		return nil, errors.Wrap(err, "bot.SendReportMessage")
	}

	err = c.blockerStorageFacade.ReportUserTx(ctx, block)
	if err != nil {
		return nil, errors.Wrap(err, "blockerStorageFacade.ReportUserTx")
	}

	return &dto.ReportUserV1Response{}, nil
}
