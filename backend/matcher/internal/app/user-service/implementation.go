package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

type userServiceUseCase interface {
	CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error)
	UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error)
	DeleteUser(ctx context.Context, in *dto.DeleteUserV1Request) (*dto.DeleteUserV1Response, error)
	GetUser(ctx context.Context, in *dto.GetUserV1Request) (*dto.GetUserV1Response, error)
	GetUsers(ctx context.Context, in *dto.GetUsersV1Request) (*dto.GetUsersV1Response, error)
}

type Implementation struct {
	desc.UnimplementedUserServiceServer
	usecase userServiceUseCase
	logger  log.Logger
}

func NewImplementation(
	logger log.Logger,
	usecase userServiceUseCase,
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
	return desc.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterUserServiceServer(gRPC, i)
}
