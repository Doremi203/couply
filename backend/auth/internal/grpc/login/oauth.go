package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
)

func (s *grpcService) OAuthLoginV1(
	ctx context.Context,
	req *login.OAuthLoginV1Request,
) (*login.OAuthLoginV1Response, error) {
	t, err := s.loginUseCase.OAuthV1(ctx, oauth.Request{
		Provider:    oauth.Provider(req.GetProvider()),
		AccessToken: oauth.Token(req.GetAccessToken()),
	})
	if err != nil {
		return nil, err
	}

	return &login.OAuthLoginV1Response{
		Token:     t.SignedString(),
		ExpiresIn: int32(t.ExpiresIn().Seconds()),
	}, nil
}
