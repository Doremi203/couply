package telegram

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotClient struct {
	api         *tgbotapi.BotAPI
	adminChatID int64
	handlers    map[string]telegram.CallbackHandler
	logger      log.Logger
}

func NewBotClient(token string, chatID int64, logger log.Logger) (*BotClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(err, "tgbotapi.NewBotAPI")
	}

	return &BotClient{
		api:         bot,
		adminChatID: chatID,
		handlers:    make(map[string]telegram.CallbackHandler),
		logger:      logger,
	}, nil
}

func (b *BotClient) RegisterCallbackHandler(prefix string, handler telegram.CallbackHandler) {
	if b.handlers == nil {
		b.handlers = make(map[string]telegram.CallbackHandler)
	}
	b.handlers[prefix] = handler
}
