package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserRequest) (*dto.ReportUserResponse, error) {
	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.GetUserID())
	if err != nil {
		return nil, err
	}

	err = c.bot.SendReportMessage(reportedUser, in.GetReportReasons(), in.GetMessage())
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserResponse{}, nil
}
