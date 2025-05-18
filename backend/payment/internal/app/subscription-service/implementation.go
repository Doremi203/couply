package subscription_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	desc "github.com/Doremi203/couply/backend/payment/gen/api/subscription-service/v1"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/subscription-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type subscriptionServiceUseCase interface {
	CreateSubscription(ctx context.Context, in *dto.CreateSubscriptionV1Request) (*dto.CreateSubscriptionV1Response, error)
	GetActiveSubscription(ctx context.Context, in *dto.GetActiveSubscriptionV1Request) (*dto.GetActiveSubscriptionV1Response, error)
}

type Implementation struct {
	desc.UnimplementedSubscriptionServiceServer
	usecase subscriptionServiceUseCase
	logger  log.Logger
}

func NewImplementation(
	logger log.Logger,
	usecase subscriptionServiceUseCase,
) *Implementation {
	return &Implementation{
		logger:  logger,
		usecase: usecase,
	}
}

func (i *Implementation) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return desc.RegisterSubscriptionServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterSubscriptionServiceServer(gRPC, i)
}
