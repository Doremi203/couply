package payment_service

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	desc "github.com/Doremi203/couply/backend/payment/gen/api/payment-service/v1"
	dto "github.com/Doremi203/couply/backend/payment/internal/dto/payment-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type paymentServiceUseCase interface {
	CreatePayment(ctx context.Context, in *dto.CreatePaymentV1Request) (*dto.CreatePaymentV1Response, error)
	GetPaymentStatus(ctx context.Context, in *dto.GetPaymentStatusV1Request) (*dto.GetPaymentStatusV1Response, error)
}

type Implementation struct {
	desc.UnimplementedPaymentServiceServer
	usecase paymentServiceUseCase
	logger  log.Logger
}

func NewImplementation(
	logger log.Logger,
	usecase paymentServiceUseCase,
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
	return desc.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterPaymentServiceServer(gRPC, i)
}
