package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotClient struct {
	api         *tgbotapi.BotAPI
	adminChatID int64
	handlers    map[string]CallbackHandler
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

func (b *BotClient) RegisterCallbackHandler(prefix string, handler CallbackHandler) {
	if b.handlers == nil {
		b.handlers = make(map[string]CallbackHandler)
	}
	b.handlers[prefix] = handler
}
