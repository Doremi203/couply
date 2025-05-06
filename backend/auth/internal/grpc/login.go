package grpc

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewLoginService(
	loginUseCase usecase.Login,
	logger log.Logger,
) *loginService {
	return &loginService{
		loginUseCase: loginUseCase,
		logger:       logger,
	}
}

type loginService struct {
	loginUseCase usecase.Login

	logger log.Logger
	login.UnimplementedLoginServer
}

func (s *loginService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return login.RegisterLoginHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *loginService) RegisterToServer(gRPC *grpc.Server) {
	login.RegisterLoginServer(gRPC, s)
}

func (s *loginService) BasicLoginV1(
	ctx context.Context,
	req *login.BasicLoginRequestV1,
) (*login.BasicLoginResponseV1, error) {
	t, err := s.loginUseCase.BasicV1(ctx, user.Email(req.GetEmail()), pswrd.Password(req.GetPassword()))
	switch {
	case errors.Is(err, usecase.ErrInvalidCredentials):
		return nil, status.Error(codes.Unauthenticated, "invalid password or username")
	case err != nil:
		s.logger.Error(errors.Wrap(err, "basic login v1 failed"))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &login.BasicLoginResponseV1{
		Token:     t.SignedString(),
		ExpiresIn: int32(t.ExpiresIn().Seconds()),
	}, nil
}
