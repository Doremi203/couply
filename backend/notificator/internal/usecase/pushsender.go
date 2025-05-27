package usecase

import (
	"context"
	"encoding/json"
	"io"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/SherClockHolmes/webpush-go"
)

func NewPushSender(
	webPushOptions *webpush.Options,
	pushRepo push.Repo,
	logger log.Logger,
) PushSender {
	return PushSender{
		webPushOptions: webPushOptions,
		pushRepo:       pushRepo,
		logger:         logger,
	}
}

type PushSender struct {
	webPushOptions *webpush.Options
	pushRepo       push.Repo
	logger         log.Logger
}

func (u PushSender) Send(ctx context.Context, r push.Recipient, p push.Push) error {
	pushJson, err := json.Marshal(p)
	if err != nil {
		return errors.WrapFailf(
			err,
			"marshal %v",
			errors.Token("push", p),
		)
	}

	for _, sub := range r.Subscriptions {
		var body []byte
		resp, err := webpush.SendNotificationWithContext(ctx, pushJson, &webpush.Subscription{
			Endpoint: string(sub.Endpoint),
			Keys: webpush.Keys{
				Auth:   sub.Credentials.AuthKey,
				P256dh: sub.Credentials.P256dh,
			},
		}, u.webPushOptions)
		if err != nil {
			u.logger.Error(errors.WrapFailf(
				err,
				"send push to %v",
				errors.Token("endpoint", sub.Endpoint),
			))
		} else {
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				u.logger.Error(errors.WrapFail(err, "read response body"))
			}
			err = resp.Body.Close()
			if err != nil {
				u.logger.Warn(errors.WrapFailf(err, "close response body"))
			}
		}

		switch resp.StatusCode {
		case 201:
			// Успех
		case 410, 404:
			err = u.pushRepo.DeleteSubscription(ctx, sub)
			if err != nil {
				u.logger.Error(errors.WrapFailf(
					err,
					"delete stale push subscription for %v",
					errors.Token("endpoint", sub.Endpoint),
				))
			}
		default:
			u.logger.Error(errors.Errorf(
				"got unexpected %v sending push for %v %v",
				errors.Token("status_code", resp.StatusCode),
				errors.Token("endpoint", sub.Endpoint),
				errors.Token("body", string(body)),
			))
		}
	}

	return nil
}
