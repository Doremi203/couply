package login

import (
	"context"

	logingprc "github.com/Doremi203/couply/backend/auth/gen/api/login"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase/login"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcService) BasicLoginV1(
	ctx context.Context,
	req *logingprc.BasicLoginRequestV1,
) (*logingprc.BasicLoginResponseV1, error) {
	t, err := s.loginUseCase.BasicV1(ctx, user.Email(req.GetEmail()), pswrd.Password(req.GetPassword()))
	switch {
	case errors.Is(err, login.ErrUserNotRegistered):
		return nil, status.Errorf(codes.NotFound, "user with %s not registered", req.GetEmail())
	case errors.Is(err, login.ErrInvalidCredentials):
		return nil, status.Error(codes.Unauthenticated, "invalid password or email")
	case err != nil:
		return nil, errors.WrapFail(err, "login with email and password")
	}

	return &logingprc.BasicLoginResponseV1{
		Token:     t.SignedString(),
		ExpiresIn: int32(t.ExpiresIn().Seconds()),
	}, nil
}
