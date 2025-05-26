package telegram

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	ApiKey   string `secret:"telegram-gateway-api-key"`
	Disabled bool
}

func NewSender(cfg Config) *sender {
	client := resty.New().
		SetBaseURL("https://gatewayapi.telegram.org/").
		SetHeader("Authorization", "Bearer "+cfg.ApiKey).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second).
		SetTimeout(3 * time.Second)

	return &sender{
		cfg:    cfg,
		client: client,
	}
}

type sender struct {
	cfg    Config
	client *resty.Client
}

func (s *sender) Send(ctx context.Context, logger log.Logger, code phoneconfirm.Code, phoneE164 user.Phone) error {
	if s.cfg.Disabled {
		return errors.Error("telegram code sender is disabled")
	}

	requestID, err := s.checkSendAbility(ctx, phoneE164)
	if err != nil {
		logger.Warn(errors.WrapFail(err, "check telegram code send ability"))
		return phoneconfirm.ErrUnsupportedPhone
	}

	var response telegramResponse[deliveryStatus]
	resp, err := s.client.R().
		SetContext(ctx).
		SetResult(&response).
		SetFormData(map[string]string{
			"request_id":   requestID,
			"phone_number": string(phoneE164),
			"code":         string(code.Value()),
			"payload":      "Your Couply security code!",
			"ttl":          "45",
		}).
		Post("sendVerificationMessage")
	if err != nil {
		return errors.WrapFail(err, "send telegram code")
	}
	if !resp.IsSuccess() {
		return errors.Errorf(
			"got unxpected %v: %v",
			errors.Token("response_status", resp.StatusCode()),
			errors.Token("body", resp.String()),
		)
	}
	if !response.Ok {
		return errors.Errorf(
			"unsuccessful telegram response: %v",
			errors.Token("error", response.Error),
		)
	}

	return nil
}

func (s *sender) checkSendAbility(ctx context.Context, phoneE164 user.Phone) (string, error) {
	var response telegramResponse[requestStatus]
	resp, err := s.client.R().
		SetContext(ctx).
		SetResult(&response).
		SetFormData(map[string]string{
			"phone_number": string(phoneE164),
		}).
		Post("checkSendAbility")
	if err != nil {
		return "", errors.WrapFail(err, "send check telegram code send ability request")
	}
	if !resp.IsSuccess() {
		return "", errors.Errorf(
			"got unxpected %v: %v",
			errors.Token("response_status", resp.StatusCode()),
			errors.Token("body", resp.String()),
		)
	}
	if !response.Ok {
		return "", errors.Errorf(
			"unsuccessful telegram response: %v",
			errors.Token("error", response.Error),
		)
	}

	return response.Result.RequestID, nil
}

type telegramResponse[T any] struct {
	Ok     bool   `json:"ok"`
	Result T      `json:"result"`
	Error  string `json:"error"`
}

type requestStatus struct {
	RequestID        string  `json:"request_id"`
	PhoneNumber      string  `json:"phone_number"`
	RequestCost      float64 `json:"request_cost"`
	IsRefunded       bool    `json:"is_refunded"`
	RemainingBalance float64 `json:"remaining_balance"`
	Payload          string  `json:"payload"`
}

type deliveryStatus struct {
	Status    string `json:"status"`
	UpdatedAt int64  `json:"updated_at"`
}
