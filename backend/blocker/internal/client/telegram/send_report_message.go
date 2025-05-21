package telegram

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"

	user_service "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *BotClient) SendReportMessage(user *user_service.User, reasons []blocker.ReportReason, message string, blockID uuid.UUID) error {
	createdAt := user.GetCreatedAt().AsTime().Format("02.01.2006 15:04")

	var reasonsText strings.Builder
	if len(reasons) > 0 {
		for i, reason := range reasons {
			reasonStr := reason.String()
			reasonStr = strings.TrimPrefix(reasonStr, "REASON_")
			reasonsText.WriteString(fmt.Sprintf("%d. %s\n", i+1, reasonStr))
		}
	} else {
		reasonsText.WriteString("не указаны")
	}

	var photosText strings.Builder
	for i, photo := range user.GetPhotos() {
		photosText.WriteString(fmt.Sprintf("%d. %s\n", i+1, photo.GetUrl()))
	}
	photosStr := photosText.String()
	if photosStr == "" {
		photosStr = "нет фото"
	}

	text := fmt.Sprintf(
		"🚨 *НОВАЯ ЖАЛОБА НА ПОЛЬЗОВАТЕЛЯ*\n\n"+
			"ID: %s\n"+
			"Имя: %s\n"+
			"Возраст: %d\n"+
			"Пол: %s\n"+
			"О себе: %s\n"+
			"Верифицирован: %t\n"+
			"Премиум: %t\n"+
			"Заблокирован: %t\n"+
			"Фото: %v\n"+
			"Аккаунт создан: %s\n\n"+
			"Причины жалобы: %v\n"+
			"Сообщение жалобы: %s",
		user.GetId(),
		user.GetName(),
		user.GetAge(),
		getGenderStr(user.GetGender()),
		user.GetBio(),
		user.GetIsVerified(),
		user.GetIsPremium(),
		user.GetIsBlocked(),
		photosStr,
		createdAt,
		reasonsText.String(),
		message,
	)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⛔ Блокировать", "block_"+blockID.String()),
			tgbotapi.NewInlineKeyboardButtonData("✅ Отклонить", "dismiss_"+blockID.String()),
		),
	)

	msg := tgbotapi.NewMessage(b.adminChatID, escapeMarkdown(text))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"

	_, err := b.api.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func getGenderStr(gender user_service.Gender) string {
	switch gender {
	case user_service.Gender_GENDER_UNSPECIFIED:
		return "Нет гендера"
	case user_service.Gender_GENDER_MALE:
		return "Мужчина"
	case user_service.Gender_GENDER_FEMALE:
		return "Женщина"
	default:
		return "Нет гендера"
	}
}

func escapeMarkdown(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(text)
}
