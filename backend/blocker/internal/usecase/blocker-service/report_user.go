package blocker_service

import (
	"context"

	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Doremi203/couply/backend/blocker/internal/dto"
)

func (c *UseCase) ReportUser(ctx context.Context, in *dto.ReportUserRequest) (*dto.ReportUserResponse, error) {
	incomingMD, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata in context")
	}

	tokenValues := incomingMD.Get("user-token")
	if len(tokenValues) == 0 {
		return nil, status.Error(codes.Unauthenticated, "user-token header is required")
	}
	userToken := tokenValues[0]

	outgoingMD := metadata.New(map[string]string{
		"user-token": userToken,
	})
	ctx = metadata.NewOutgoingContext(ctx, outgoingMD)

	user, err := c.userServiceClient.GetUserV1(ctx)
	if err != nil {
		return nil, err
	}

	c.tokenCache.Set(user.GetId(), userToken, cache.DefaultExpiration)

	err = c.bot.SendReportMessage(user, in.GetReportReasons(), in.GetMessage())
	if err != nil {
		return nil, err
	}

	return &dto.ReportUserResponse{}, nil
}
