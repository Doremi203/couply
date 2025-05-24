package oauth

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/go-resty/resty/v2"
)

func newCodeExchanger(
	clientID string,
	authURL string,
	redirectURL string,
) codeExchanger {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)
	return codeExchanger{
		clientID:    clientID,
		client:      client,
		authURL:     authURL,
		redirectURL: redirectURL,
	}
}

type codeExchanger struct {
	clientID    string
	authURL     string
	redirectURL string
	client      *resty.Client
}

func (p codeExchanger) ExchangeCodeForToken(
	ctx context.Context,
	code Code,
	state State,
	codeVerifier CodeVerifier,
	deviceID DeviceID,
) (Token, error) {
	var exchangeCodeResponse exchangeCodeResponse
	var exchangeCodeError exchangeCodeError

	formData := map[string]string{
		"grant_type":    "authorization_code",
		"code":          string(code),
		"code_verifier": string(codeVerifier),
		"client_id":     p.clientID,
		"redirect_uri":  p.redirectURL,
		"state":         string(state),
		"device_id":     string(deviceID),
	}

	resp, err := p.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(&exchangeCodeResponse).
		SetError(&exchangeCodeError).
		Post(p.authURL)
	if err != nil {
		return "", errors.WrapFail(err, "exchange code for token")
	}
	if !resp.IsSuccess() || exchangeCodeResponse.Error != "" {
		var errorMsg string
		var errorDescription string
		if exchangeCodeResponse.Error != "" {
			errorMsg = exchangeCodeResponse.Error
			errorDescription = exchangeCodeResponse.ErrorDescription
		} else {
			errorMsg = exchangeCodeError.Error
			errorDescription = exchangeCodeError.ErrorDescription
		}
		return "", errors.Errorf(
			"got unsuccessful %v from %v: %v %v",
			errors.Token("status_code", resp.StatusCode()),
			errors.Token("api", p.authURL),
			errors.Token("error", errorMsg),
			errors.Token("error_description", errorDescription),
		)
	}

	return Token(exchangeCodeResponse.AccessToken), nil
}

type exchangeCodeResponse struct {
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

type exchangeCodeError struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}
