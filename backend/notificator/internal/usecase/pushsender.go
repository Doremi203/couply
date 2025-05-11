package usecase

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/SherClockHolmes/webpush-go"
)

type PushSender struct {
	webPushOptions *webpush.Options
	pushRepo       push.Repo
	logger         log.Logger
}

func (u PushSender) Send(ctx context.Context, r push.Recipient, p push.Push) error {
	for _, sub := range r.Subscriptions {
		resp, err := webpush.SendNotificationWithContext(ctx, []byte(p.Text), &webpush.Subscription{
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
		}
		err = resp.Body.Close()
		if err != nil {
			u.logger.Warn(errors.WrapFailf(err, "close response body"))
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
			u.logger.Error(errors.Wrapf(
				err,
				"got unexpected %v sending push for %v",
				errors.Token("status_code", resp.StatusCode),
				errors.Token("endpoint", sub.Endpoint),
			))
		}
	}

	return nil
}
