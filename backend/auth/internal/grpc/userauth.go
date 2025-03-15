package grpc

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/gen/api/registration"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	registrationUC "github.com/Doremi203/couply/backend/auth/internal/usecase/registration"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func NewRegistrationService(
	useCase registrationUC.Basic,
	log *slog.Logger,
) *userAuthService {
	return &userAuthService{
		useCase: useCase,
		log:     log,
	}
}

type userAuthService struct {
	useCase registrationUC.Basic

	log *slog.Logger
	registration.UnimplementedRegistrationServer
}

func (s *userAuthService) RegisterToServer(gRPC *grpc.Server) {
	registration.RegisterRegistrationServer(gRPC, s)
}

func (s *userAuthService) Register(
	ctx context.Context,
	req *registration.BasicRegisterRequest,
) (*registration.BasicRegisterResponse, error) {
	err := s.useCase.Run(ctx, user.Email(req.GetEmail()), user.Password(req.GetPassword()))
	switch {
	case errors.Is(err, registrationUC.ErrAlreadyRegistered):
		return nil, status.Errorf(codes.AlreadyExists, "user with email %s %v", req.GetEmail(), err)
	case err != nil:
		s.log.Error("failed to register user", "error", err)
		return nil, status.Error(codes.Internal, "internal server error")
	default:
		// continue
	}

	return &registration.BasicRegisterResponse{}, nil
}
