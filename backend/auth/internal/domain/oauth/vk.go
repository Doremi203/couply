package oauth

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/go-resty/resty/v2"
)

type VKConfig struct {
	ClientID string `secret:"oauth-vk-client-id"`
}

func newVkProvider(cfg VKConfig) *vkProvider {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)
	return &vkProvider{
		client: client,
		codeExchanger: newCodeExchanger(
			cfg.ClientID,
			"https://id.vk.com/oauth2/auth",
			"https://testing.couply.ru",
		),
	}
}

type vkProvider struct {
	client *resty.Client

	codeExchanger
}

func (p *vkProvider) FetchUserInfo(
	ctx context.Context,
	token Token,
) (UserInfo, error) {
	var apiResp vkUserInfoResponse
	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"access_token": string(token),
			"client_id":    p.codeExchanger.clientID,
		}).
		SetResult(&apiResp).
		Post("https://id.vk.com/oauth2/user_info")
	if err != nil {
		return UserInfo{}, errors.Wrapf(
			err,
			"failed to fetch user info from %v",
			errors.Token("provider", VKProvider),
		)
	}
	if !resp.IsSuccess() {
		return UserInfo{}, errors.Errorf(
			"got unsuccessful %v from %v api",
			errors.Token("status_code", resp.StatusCode()),
			errors.Token("provider", VKProvider),
		)
	}

	if apiResp.User.UserId == "" {
		return UserInfo{}, errors.Errorf(
			"got empty provider user id from %v api",
			errors.Token("provider", VKProvider),
		)
	}
	if apiResp.User.Email == "" {
		return UserInfo{}, errors.Errorf(
			"got empty email from %v api",
			errors.Token("provider", VKProvider),
		)
	}

	return UserInfo{
		ProviderUserID: ProviderUserID(apiResp.User.UserId),
		Email:          apiResp.User.Email,
		Phone:          apiResp.User.Phone,
	}, nil
}

type vkUserInfoResponse struct {
	User struct {
		UserId    string `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Avatar    string `json:"avatar"`
		Email     string `json:"email"`
		Sex       int    `json:"sex"`
		Verified  bool   `json:"verified"`
		Birthday  string `json:"birthday"`
	} `json:"user"`
}
