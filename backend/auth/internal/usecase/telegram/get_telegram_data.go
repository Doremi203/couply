package telegram

import (
	"context"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
)

const (
	tgURLFirstPart = "https://t.me/"
)

func (u UseCase) GetTelegramData(ctx context.Context, userID string) (string, error) {
	_, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "token.GetUserIDFromContext")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return "", errors.Wrap(err, "uuid.Parse")
	}

	tgData, err := u.telegramRepo.GetTelegramData(ctx, userUUID)
	if err != nil {
		return "", errors.Wrap(err, "u.telegramRepo.SetTelegramData")
	}

	tgURL := tgURLFirstPart + tgData.TelegramUsername

	return tgURL, nil
}
