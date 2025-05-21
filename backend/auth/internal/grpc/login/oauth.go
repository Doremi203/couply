package login

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *grpcService) OAuthLoginV1(
	ctx context.Context,
	req *login.OAuthLoginV1Request,
) (*login.OAuthLoginV1Response, error) {
	oauthResp, err := s.loginUseCase.OAuthV1(ctx, oauth.Request{
		Provider:    oauth.ProviderType(req.GetProvider()),
		AccessToken: oauth.Token(req.GetAccessToken()),
		Code:        oauth.Code(req.GetCode()),
		State:       oauth.State(req.GetState()),
	})
	if err != nil {
		return nil, err
	}

	path := "home"
	if oauthResp.IsFirstLogin {
		path = "enterInfo"
	}

	redirectURL := fmt.Sprintf(
		"%s/%s",
		s.config.OAuthRedirectFrontendDomain,
		path,
	)

	accessTokenCookie := fmt.Sprintf("access_token=%s; Path=/; HttpOnly; SameSite=Strict; Max-Age=%d",
		oauthResp.TokenPair.AccessToken.SignedString(),
		int(oauthResp.TokenPair.AccessToken.ExpiresIn().Seconds()))

	refreshTokenCookie := fmt.Sprintf("refresh_token=%s; Path=/; HttpOnly; SameSite=Strict; Max-Age=%d",
		string(oauthResp.TokenPair.RefreshToken.Token),
		int(oauthResp.TokenPair.RefreshToken.ExpiresIn.Seconds()))

	md := metadata.Pairs(
		"Location", redirectURL,
		"Set-Cookie", accessTokenCookie,
		"Set-Cookie", refreshTokenCookie,
	)
	err = grpc.SendHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	return &login.OAuthLoginV1Response{}, nil
}
