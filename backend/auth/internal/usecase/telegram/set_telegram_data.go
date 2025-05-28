package telegram

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
)

func (u UseCase) SetTelegramData(ctx context.Context, _ int64, _ string, _ string, id int64, _ string, username string) (telegram.Data, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return telegram.Data{}, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	tgData := telegram.Data{
		UserID:           userID,
		TelegramID:       id,
		TelegramUsername: username,
	}

	err = u.telegramRepo.UpsertTelegramData(ctx, tgData)
	if err != nil {
		return telegram.Data{}, errors.Wrap(err, "u.telegramRepo.SetTelegramData")
	}

	return tgData, nil
}
