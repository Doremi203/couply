package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotClient struct {
	api         *tgbotapi.BotAPI
	adminChatID int64
}

func NewBotClient(token string, chatID int64) (*BotClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &BotClient{
		api:         bot,
		adminChatID: chatID,
	}, nil
}
