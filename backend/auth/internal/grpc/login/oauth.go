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
		"%s/%s#token=%s&expiresAt=%d",
		s.config.OAuthRedirectFrontendDomain,
		path,
		oauthResp.Token.SignedString(),
		int(oauthResp.Token.ExpiresIn().Seconds()),
	)

	md := metadata.Pairs(
		"Location", redirectURL,
	)
	err = grpc.SendHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	return &login.OAuthLoginV1Response{}, nil
}
