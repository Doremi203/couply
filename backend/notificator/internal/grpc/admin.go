package grpc

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	pushgrpc "github.com/Doremi203/couply/backend/notificator/gen/api/push"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/Doremi203/couply/backend/notificator/internal/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewAdminService(
	pushRepo push.Repo,
	pushSender usecase.PushSender,
	logger log.Logger,
) *adminService {
	return &adminService{
		pushRepo:   pushRepo,
		pushSender: pushSender,
		logger:     logger,
	}
}

type adminService struct {
	pushRepo   push.Repo
	pushSender usecase.PushSender

	logger log.Logger
	pushgrpc.UnimplementedAdminServer
}

func (s *adminService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return pushgrpc.RegisterAdminHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *adminService) RegisterToServer(gRPC *grpc.Server) {
	pushgrpc.RegisterAdminServer(gRPC, s)
}

func (s *adminService) SendPushV1(
	ctx context.Context,
	req *pushgrpc.SendPushV1Request,
) (*pushgrpc.SendPushV1Response, error) {
	recipients, err := s.pushRepo.GetAllRecipients(ctx)
	if err != nil {
		return nil, err
	}

	for _, recipient := range recipients {
		err = s.pushSender.Send(ctx, recipient, push.Push{
			Text: req.GetText(),
		})
		if err != nil {
			s.logger.Error(errors.Errorf(
				"failed to send push notification to %v", errors.Token(
					"recipient_id",
					recipient.ID,
				)),
			)
			continue
		}
	}

	return &pushgrpc.SendPushV1Response{}, nil
}
