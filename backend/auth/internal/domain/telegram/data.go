package telegram

import "github.com/google/uuid"

type Data struct {
	UserID           uuid.UUID `json:"user_id"`
	TelegramID       string    `json:"telegram_id"`
	TelegramUsername string    `json:"telegram_username"`
}
