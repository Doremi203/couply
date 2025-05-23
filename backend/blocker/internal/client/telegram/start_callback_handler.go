package telegram

import (
	"fmt"
	"strings"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	callbackQueryTimeout = 60
	callbackDataParts    = 2
)

func (b *BotClient) StartCallbackHandler() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = callbackQueryTimeout
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery == nil {
			continue
		}

		data := update.CallbackQuery.Data
		msg := update.CallbackQuery.Message

		prefix, blockID, ok := b.parseCallbackData(data)
		if !ok {
			b.logger.Warn(errors.Errorf("invalid callback format: %s", data))
			continue
		}

		handler, exists := b.handlers[prefix]
		if !exists {
			b.logger.Warn(errors.Errorf("no handler for prefix: %s", prefix))
			continue
		}

		result := handler(blockID)
		b.handleCallbackResult(update, msg, result)
	}
}

func (b *BotClient) parseCallbackData(data string) (string, string, bool) {
	parts := strings.SplitN(data, "_", callbackDataParts)
	if len(parts) != callbackDataParts {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func (b *BotClient) handleCallbackResult(update tgbotapi.Update, msg *tgbotapi.Message, result telegram.CallbackResult) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Обработано")
	if _, err := b.api.Request(callback); err != nil {
		b.logger.Error(errors.Wrap(err, "failed to acknowledge callback"))
	}

	responseText := b.prepareResponseText(result)
	actionText := b.prepareActionText(result)

	b.updateOriginalMessage(msg, actionText)
	b.sendResponseMessage(msg.Chat.ID, responseText)
}

func (b *BotClient) prepareResponseText(result telegram.CallbackResult) string {
	if result.Error != nil {
		return fmt.Sprintf("❌ Ошибка: %s", result.ResponseText)
	}
	return result.ResponseText
}

func (b *BotClient) prepareActionText(result telegram.CallbackResult) string {
	if result.Error != nil {
		return "❌ Ошибка"
	}
	return result.ActionText
}

func (b *BotClient) updateOriginalMessage(msg *tgbotapi.Message, actionText string) {
	editedText := fmt.Sprintf("%s\n\n%s", msg.Text, actionText)
	editMsg := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, editedText)
	editMsg.ParseMode = tgbotapi.ModeMarkdown

	if _, err := b.api.Send(editMsg); err != nil {
		b.logger.Error(errors.Wrapf(err, "failed to edit message with chat_id %v and message_id %v", msg.Chat.ID, msg.MessageID))
	}

	b.removeInlineKeyboard(msg)
}

func (b *BotClient) removeInlineKeyboard(msg *tgbotapi.Message) {
	editMarkup := tgbotapi.NewEditMessageReplyMarkup(
		msg.Chat.ID,
		msg.MessageID,
		tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}},
	)

	if _, err := b.api.Send(editMarkup); err != nil {
		b.logger.Error(errors.Wrapf(err, "failed to remove inline keyboard with chat_id %v and message_id %v", msg.Chat.ID, msg.MessageID))
	}
}

func (b *BotClient) sendResponseMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := b.api.Send(msg); err != nil {
		b.logger.Error(errors.Wrapf(err, "failed to send response message with chat_id %v and text %v", chatID, text))

	}
}
