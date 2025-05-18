package oauth

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/go-resty/resty/v2"
)

func newYandexFetcher() *yandexFetcher {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)
	return &yandexFetcher{client: client}
}

type yandexFetcher struct {
	client *resty.Client
}

func (f *yandexFetcher) Fetch(ctx context.Context, token Token) (UserInfo, error) {
	var apiResp apiResponse
	resp, err := f.client.R().
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

type apiResponse struct {
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
