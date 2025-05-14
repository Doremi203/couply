package telegram

import (
	"encoding/base64"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *BotClient) StartCallbackHandler(blockCallback func(userID string, userToken string)) {
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
				encodedData := strings.TrimPrefix(data, "block_")
				decodedBytes, err := base64.URLEncoding.DecodeString(encodedData)
				if err != nil {
					log.Printf("Failed to decode data: %v", err)
					continue
				}
				parts := strings.Split(string(decodedBytes), "|")
				if len(parts) != 2 {
					log.Printf("Invalid data format: %s", string(decodedBytes))
					continue
				}
				userID := parts[0]
				userToken := parts[1]
				blockCallback(userID, userToken)
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
