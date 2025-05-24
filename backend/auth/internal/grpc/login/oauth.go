package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	tokenPb "github.com/Doremi203/couply/backend/auth/gen/api/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
)

func (s *grpcService) OAuthLoginV1(
	ctx context.Context,
	req *login.OAuthLoginV1Request,
) (*login.OAuthLoginV1Response, error) {
	oauthResp, err := s.loginUseCase.OAuthV1(ctx, oauth.Request{
		Provider:     oauth.ProviderType(req.GetProvider()),
		DeviceID:     oauth.DeviceID(req.GetDeviceId()),
		Code:         oauth.Code(req.GetCode()),
		State:        oauth.State(req.GetState()),
		CodeVerifier: oauth.CodeVerifier(req.GetCodeVerifier()),
	})
	if err != nil {
		return nil, err
	}

	return &login.OAuthLoginV1Response{
		AccessToken: &tokenPb.Token{
			Token:     oauthResp.TokenPair.AccessToken.SignedString(),
			ExpiresIn: int32(oauthResp.TokenPair.AccessToken.ExpiresIn().Seconds()),
		},
		RefreshToken: &tokenPb.Token{
			Token:     string(oauthResp.TokenPair.RefreshToken.Token),
			ExpiresIn: int32(oauthResp.TokenPair.RefreshToken.ExpiresIn.Seconds()),
		},
		FirstLogin: oauthResp.IsFirstLogin,
	}, nil
}
