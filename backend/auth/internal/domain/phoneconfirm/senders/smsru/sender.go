package smsru

import (
	"context"
	"strings"

	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	ApiKey string `secret:"smsru-api-key"`
	Test   bool
}

func NewSender(
	cfg Config,
	httpClient *resty.Client,
) *sender {
	return &sender{
		cfg:        cfg,
		httpClient: httpClient,
	}
}

type sender struct {
	cfg        Config
	httpClient *resty.Client
}

func (s *sender) Send(ctx context.Context, logger log.Logger, code phoneconfirm.Code, phoneE164 user.Phone) error {
	toParam := strings.TrimPrefix(string(phoneE164), "+")

	var result struct {
		Status     string `json:"status"`
		StatusCode int    `json:"status_code"`
		StatusText string `json:"status_text"`
		Sms        map[string]struct {
			Status     string `json:"status"`
			StatusCode int    `json:"status_code"`
			SmsID      string `json:"sms_id"`
			StatusText string `json:"status_text,omitempty"`
		} `json:"sms"`
		Balance float64 `json:"balance"`
	}

	query := map[string]string{
		"api_id": s.cfg.ApiKey,
		"to":     toParam,
		"msg":    string(code.Value()),
		"json":   "1",
	}
	if s.cfg.Test {
		query["test"] = "1"
		logger.Infof(
			"Sending %v to %v",
			errors.Token("code", code.Value()),
			errors.Token("phone", phoneE164),
		)
	}

	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetQueryParams(query).
		Get("https://sms.ru/sms/send")
	if err != nil {
		return errors.WrapFail(err, "send sms")
	}
	if !resp.IsSuccess() {
		return errors.Errorf(
			"got unexpected %v for send sms",
			errors.Token("http_status_code", resp.StatusCode()),
		)
	}

	if result.StatusCode != 100 {
		return errors.Errorf(
			"got sms ru api error %v %v",
			errors.Token("status_code", result.StatusCode),
			errors.Token("details", result.StatusText),
		)
	}

	for _, info := range result.Sms {
		switch info.StatusCode {
		case 100:
			continue
		case 204:
			return phoneconfirm.ErrUnsupportedPhone
		default:
			return errors.Errorf(
				"got sms ru api error for sms %v %v",
				errors.Token("status_code", info.StatusCode),
				errors.Token("details", info.StatusText),
			)
		}
	}

	return nil
}
