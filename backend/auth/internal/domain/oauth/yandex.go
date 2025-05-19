package oauth

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/go-resty/resty/v2"
)

type YandexConfig struct {
	ClientID     string `secret:"oauth-yandex-client-id"`
	ClientSecret string `secret:"oauth-yandex-client-secret"`
}

func newYandexProvider(config YandexConfig) *yandexProvider {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)
	return &yandexProvider{
		config: config,
		client: client,
	}
}

type yandexProvider struct {
	config YandexConfig
	client *resty.Client
}

func (p *yandexProvider) ExchangeCodeForToken(ctx context.Context, code Code, state State) (Token, error) {
	var accessTokenResponse accessTokenOKResponse
	var errorResponse accessTokenErrorResponse

	formData := map[string]string{
		"grant_type":    "authorization_code",
		"code":          string(code),
		"client_id":     p.config.ClientID,
		"client_secret": p.config.ClientSecret,
	}
	if state != "" {
		formData["state"] = string(state)
	}

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&accessTokenResponse).
		SetError(&errorResponse).
		Post("https://oauth.yandex.ru/token")
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to exchange code for token from %v",
			errors.Token("provider", YandexProvider),
		)
	}
	if !resp.IsSuccess() {
		return "", errors.Errorf(
			"got unsuccessful %v from %v api: %v %v",
			errors.Token("status_code", resp.StatusCode()),
			errors.Token("provider", YandexProvider),
			errors.Token("error", errorResponse.Error),
			errors.Token("error_description", errorResponse.ErrorDescription),
		)
	}

	return Token(accessTokenResponse.AccessToken), nil
}

type accessTokenOKResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type accessTokenErrorResponse struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

func (p *yandexProvider) FetchUserInfo(ctx context.Context, token Token) (UserInfo, error) {
	var apiResp userInfoResponse
	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader("Authorization", fmt.Sprintf("OAuth %s", token)).
		SetResult(&apiResp).
		Get("https://login.yandex.ru/info")
	if err != nil {
		return UserInfo{}, errors.Wrapf(
			err,
			"failed to fetch user info from %v",
			errors.Token("provider", YandexProvider),
		)
	}
	if !resp.IsSuccess() {
		return UserInfo{}, errors.Errorf(
			"got unsuccessful %v from %v api",
			errors.Token("status_code", resp.StatusCode()),
			errors.Token("provider", YandexProvider),
		)
	}

	if apiResp.Id == "" {
		return UserInfo{}, errors.Errorf(
			"got empty provider user id from %v api",
			errors.Token("provider", YandexProvider),
		)
	}
	if apiResp.DefaultEmail == "" {
		return UserInfo{}, errors.Errorf(
			"got empty email from %v api",
			errors.Token("provider", YandexProvider),
		)
	}
	if apiResp.DefaultPhone.Number == "" {
		return UserInfo{}, errors.Errorf(
			"got empty phone from %v api",
			errors.Token("provider", YandexProvider),
		)
	}

	return UserInfo{
		ProviderUserID: ProviderUserID(apiResp.Id),
		Email:          apiResp.DefaultEmail,
		Phone:          apiResp.DefaultPhone.Number,
	}, nil
}

type userInfoResponse struct {
	Id           string   `json:"id"`
	Login        string   `json:"login"`
	ClientId     string   `json:"client_id"`
	DisplayName  string   `json:"display_name"`
	RealName     string   `json:"real_name"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Sex          string   `json:"sex"`
	DefaultEmail string   `json:"default_email"`
	Emails       []string `json:"emails"`
	DefaultPhone struct {
		Id     int    `json:"id"`
		Number string `json:"number"`
	} `json:"default_phone"`
	Psuid string `json:"psuid"`
}
