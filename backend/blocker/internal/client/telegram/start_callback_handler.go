package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *BotClient) StartCallbackHandler(blockCallback func(userID string)) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			msg := update.CallbackQuery.Message
			userID := strings.TrimPrefix(data, "block_")
			userID = strings.TrimPrefix(userID, "dismiss_")

			var responseText string
			var actionText string

			if strings.HasPrefix(data, "block_") {
				blockCallback(userID)
				responseText = "Пользователь был успешно заблокирован"
				actionText = "⛔ Заблокировано"
			} else {
				responseText = "Жалоба отклонена"
				actionText = "✅ Отклонено"
			}

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Обработано")
			if _, err := b.api.Request(callback); err != nil {
				log.Println(err)
				continue
			}

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

			replyMsg := tgbotapi.NewMessage(msg.Chat.ID, responseText)
			if _, err := b.api.Send(replyMsg); err != nil {
				log.Println(err)
			}
		}
	}
}
