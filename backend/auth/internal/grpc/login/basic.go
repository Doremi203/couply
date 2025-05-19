package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	login2 "github.com/Doremi203/couply/backend/auth/internal/usecase/login"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcService) BasicLoginV1(
	ctx context.Context,
	req *login.BasicLoginRequestV1,
) (*login.BasicLoginResponseV1, error) {
	t, err := s.loginUseCase.BasicV1(ctx, user.Email(req.GetEmail()), pswrd.Password(req.GetPassword()))
	switch {
	case errors.Is(err, login2.ErrInvalidCredentials):
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
