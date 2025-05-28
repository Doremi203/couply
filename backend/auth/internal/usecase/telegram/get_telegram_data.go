package telegram

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
)

func (u UseCase) GetTelegramData(ctx context.Context) (string, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "token.GetUserIDFromContext")
	}

	tgData, err := u.telegramRepo.GetTelegramData(ctx, userID)
	if err != nil {
		return "", errors.Wrap(err, "u.telegramRepo.SetTelegramData")
	}

	tgURL := "https://t.me/" + tgData.TelegramUsername

	return tgURL, nil
}
