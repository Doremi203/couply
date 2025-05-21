package token

import (
	"context"

	tokenPb "github.com/Doremi203/couply/backend/auth/gen/api/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	tokenUC "github.com/Doremi203/couply/backend/auth/internal/usecase/token"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcService) Refresh(
	ctx context.Context,
	req *tokenPb.RefreshRequest,
) (*tokenPb.RefreshResponse, error) {
	refreshResponse, err := s.useCase.Refresh(ctx, token.RefreshValue(req.GetRefreshToken()))
	switch {
	case errors.Is(err, tokenUC.ErrInvalidRefreshToken):
		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
	case err != nil:
		return nil, errors.WrapFail(err, "refresh token")
	}

	return &tokenPb.RefreshResponse{
		AccessToken: &tokenPb.Token{
			Token:     refreshResponse.AccessToken.SignedString(),
			ExpiresIn: int32(refreshResponse.AccessToken.ExpiresIn().Seconds()),
		},
		RefreshToken: &tokenPb.Token{
			Token:     string(refreshResponse.RefreshToken.Token),
			ExpiresIn: int32(refreshResponse.RefreshToken.ExpiresIn.Seconds()),
		},
	}, nil
}
