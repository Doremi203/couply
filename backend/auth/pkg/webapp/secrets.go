package webapp

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
)

func (a *App) loadSecrets() error {
	for _, id := range a.Config.secrets.Ids {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		secret, err := a.ycSDKClient.LockboxPayload().Payload().Get(ctx, &lockbox.GetPayloadRequest{
			SecretId: id,
		})
		cancel()
		if err != nil {
			return errors.WrapFail(err, "get token secret")
		}
		if len(secret.GetEntries()) == 0 {
			return errors.Error("secret required with config but not found in yc")
		}

		for _, entry := range secret.GetEntries() {
			_, ok := a.Config.secretsMap[entry.GetKey()]
			if ok {
				return errors.Errorf("all keys must be unique, non unique %v", errors.Token("secret_key", entry.GetKey()))
			}
			a.Config.secretsMap[entry.GetKey()] = entry.GetTextValue()
		}
	}

	return nil
}
