package grpc

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	pushgrpc "github.com/Doremi203/couply/backend/notificator/gen/api/push"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/Doremi203/couply/backend/notificator/internal/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewPushSubscriptionService(
	pushSubscriptionUseCase usecase.PushSubscription,
	logger log.Logger,
) *pushSubscriptionService {
	return &pushSubscriptionService{
		pushSubscriptionUseCase: pushSubscriptionUseCase,
		logger:                  logger,
	}
}

type pushSubscriptionService struct {
	pushSubscriptionUseCase usecase.PushSubscription

	logger log.Logger
	pushgrpc.UnimplementedSubscriptionServer
}

func (s *pushSubscriptionService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return pushgrpc.RegisterSubscriptionHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *pushSubscriptionService) RegisterToServer(gRPC *grpc.Server) {
	pushgrpc.RegisterSubscriptionServer(gRPC, s)
}

func (s *pushSubscriptionService) SubscribeV1(
	ctx context.Context,
	req *pushgrpc.SubscribeV1Request,
) (*pushgrpc.SubscribeV1Response, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	recipientID := push.RecipientID(t.GetUserID())

	subscription, err := push.NewSubscription(
		recipientID,
		req.GetEndpoint(),
		req.GetP256Dh(),
		req.GetAuthKey(),
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.pushSubscriptionUseCase.Subscribe(ctx, subscription)
	if err != nil {
		return nil, errors.WrapFailf(err, "subscribe %v to pushes", errors.Token("user_id", recipientID))
	}

	return &pushgrpc.SubscribeV1Response{}, nil
}

func (s *pushSubscriptionService) UnsubscribeV1(
	ctx context.Context,
	req *pushgrpc.UnsubscribeV1Request,
) (*pushgrpc.UnsubscribeV1Response, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	if req.GetEndpoint() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing endpoint")
	}

	recipientID := push.RecipientID(t.GetUserID())

	err := s.pushSubscriptionUseCase.Unsubscribe(ctx, push.Subscription{
		RecipientID: recipientID,
		Endpoint:    push.Endpoint(req.GetEndpoint()),
	})
	if err != nil {
		return nil, errors.WrapFailf(err, "unsubscribe %v to pushes", errors.Token("recipient_id", recipientID))
	}

	return &pushgrpc.UnsubscribeV1Response{}, nil
}
