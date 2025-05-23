package payment_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/payments/gen/api/payment-service/v1"
	dto "github.com/Doremi203/couply/backend/payments/internal/dto/payment-service"
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
}

func NewImplementation(
	usecase paymentServiceUseCase,
) *Implementation {
	return &Implementation{
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
