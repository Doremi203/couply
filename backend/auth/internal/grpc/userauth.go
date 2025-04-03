package grpc

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/gen/api/registration"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func NewRegistrationService(
	useCase usecase.Registration,
	log *slog.Logger,
) *userAuthService {
	return &userAuthService{
		useCase: useCase,
		log:     log,
	}
}

type userAuthService struct {
	useCase usecase.Registration

	log *slog.Logger
	registration.UnimplementedRegistrationServer
}

func (s *userAuthService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return registration.RegisterRegistrationHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *userAuthService) RegisterToServer(gRPC *grpc.Server) {
	registration.RegisterRegistrationServer(gRPC, s)
}

func (s *userAuthService) BasicRegister(
	ctx context.Context,
	req *registration.BasicRegisterRequest,
) (*registration.BasicRegisterResponse, error) {
	err := s.useCase.BasicRegister(ctx, user.Email(req.GetEmail()), pswrd.Password(req.GetPassword()))
	switch {
	case errors.Is(err, usecase.ErrAlreadyRegistered):
		return nil, status.Errorf(codes.AlreadyExists, "user with email %s %v", req.GetEmail(), err)
	case err != nil:
		s.log.Error("failed to register user", "error", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &registration.BasicRegisterResponse{}, nil
}
