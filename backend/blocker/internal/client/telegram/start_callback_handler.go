package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *BotClient) StartCallbackHandler() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery == nil {
			continue
		}

		data := update.CallbackQuery.Data
		msg := update.CallbackQuery.Message

		parts := strings.SplitN(data, "_", 2)
		if len(parts) != 2 {
			log.Printf("Invalid callback format: %s", data)
			continue
		}

		prefix := parts[0]
		blockID := parts[1]

		handler, exists := b.handlers[prefix]
		if !exists {
			log.Printf("No handler for prefix: %s", prefix)
			continue
		}

		result := handler(blockID)
		b.handleCallbackResult(update, msg, result)
	}
}

func (b *BotClient) handleCallbackResult(update tgbotapi.Update, msg *tgbotapi.Message, result CallbackResult) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Обработано")
	if _, err := b.api.Request(callback); err != nil {
		log.Println(err)
	}

	responseText, actionText := b.prepareResponseTexts(result)
	b.updateOriginalMessage(msg, actionText)
	b.sendResponseMessage(msg, responseText)
}

func (b *BotClient) prepareResponseTexts(result CallbackResult) (string, string) {
	if result.Error != nil {
		return "❌ Ошибка: " + result.Error.Error(), "❌ Ошибка"
	}
	return result.ResponseText, result.ActionText
}

func (b *BotClient) updateOriginalMessage(msg *tgbotapi.Message, actionText string) {
	editedText := msg.Text + "\n\n" + actionText
	editMsg := tgbotapi.NewEditMessageText(msg.Chat.ID, msg.MessageID, editedText)
	editMsg.ParseMode = "Markdown"

	editMarkup := tgbotapi.NewEditMessageReplyMarkup(
		msg.Chat.ID,
		msg.MessageID,
		tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}},
	)

	if _, err := b.api.Send(editMsg); err != nil {
		log.Println(err)
	}
	if _, err := b.api.Send(editMarkup); err != nil {
		log.Println(err)
	}
}

func (b *BotClient) sendResponseMessage(msg *tgbotapi.Message, text string) {
	replyMsg := tgbotapi.NewMessage(msg.Chat.ID, text)
	if _, err := b.api.Send(replyMsg); err != nil {
		log.Println(err)
	}
}
