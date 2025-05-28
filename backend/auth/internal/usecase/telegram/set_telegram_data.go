package telegram

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
)

const (
	idData       = "id"
	usernameData = "username"
)

func (u UseCase) SetTelegramData(ctx context.Context, authData map[string]string) (telegram.Data, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return telegram.Data{}, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	tgData := telegram.Data{
		UserID:           userID,
		TelegramID:       authData[idData],
		TelegramUsername: authData[usernameData],
	}

	err = u.telegramRepo.UpsertTelegramData(ctx, tgData)
	if err != nil {
		return telegram.Data{}, errors.Wrap(err, "u.telegramRepo.SetTelegramData")
	}

	return tgData, nil
}
