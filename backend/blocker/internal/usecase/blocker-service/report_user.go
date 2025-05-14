package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"google.golang.org/grpc/metadata"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserRequest) (*dto.ReportUserResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	userToken := md.Get("user-token")[0]

	reportedUser, err := c.userServiceClient.GetUserByIDV1(ctx, in.GetUserID(), userToken)
	if err != nil {
		return nil, err
	}

	err = c.bot.SendReportMessage(reportedUser, in.GetReportReasons(), in.GetMessage(), userToken)
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserResponse{}, nil
}
