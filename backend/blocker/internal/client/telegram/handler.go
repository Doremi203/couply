package telegram

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/google/uuid"
)

type userClient interface {
	GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error)
	UpdateUserByIDV1(ctx context.Context, user *userservicegrpc.User) error
}

type botClient interface {
	RegisterCallbackHandler(prefix string, handler CallbackHandler)
}

type blockerStorageFacade interface {
	DeleteUserBlockTx(ctx context.Context, blockID uuid.UUID) error
	GetBlockInfoTx(ctx context.Context, userID uuid.UUID) (*blocker.UserBlock, error)
}

type BotHandler struct {
	userClient    userClient
	blockerFacade blockerStorageFacade
	bot           botClient
	logger        log.Logger
}

func NewBotHandler(
	userClient userClient,
	blockerFacade blockerStorageFacade,
	bot botClient,
	logger log.Logger,
) *BotHandler {
	return &BotHandler{
		userClient:    userClient,
		blockerFacade: blockerFacade,
		bot:           bot,
		logger:        logger,
	}
}

func (h *BotHandler) SetupRoutes() {
	h.bot.RegisterCallbackHandler("block", h.handleBlockAction)
	h.bot.RegisterCallbackHandler("dismiss", h.handleDismissAction)
}

func (h *BotHandler) handleBlockAction(callbackData string) CallbackResult {
	userID := strings.TrimPrefix(callbackData, "block_")
	return h.processBlockRequest(userID)
}

func (h *BotHandler) handleDismissAction(callbackData string) CallbackResult {
	userID := strings.TrimPrefix(callbackData, "dismiss_")
	return h.processDismissRequest(userID)
}

func (h *BotHandler) processBlockRequest(userID string) (retErr CallbackResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		retErr = CallbackResult{
			Error: fmt.Errorf("invalid user ID format: %w", err),
		}
	}

	defer h.cleanupOnError(ctx, userUUID, err)

	userData, err := h.userClient.GetUserByIDV1(ctx, userID)
	if err != nil {
		retErr = h.errorResult(
			err,
			"❌ Ошибка блокировки",
			"❌ Не удалось получить данные пользователя",
		)
	}

	userData.IsBlocked = true
	if err := h.userClient.UpdateUserByIDV1(ctx, userData); err != nil {
		retErr = h.errorResult(
			err,
			"❌ Ошибка блокировки",
			"❌ Не удалось обновить статус пользователя",
		)
	}

	return CallbackResult{
		ActionText:   "✅ Заблокирован",
		ResponseText: "Пользователь успешно заблокирован",
	}
}

func (h *BotHandler) processDismissRequest(userID string) CallbackResult {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return CallbackResult{
			Error: fmt.Errorf("invalid user ID format: %w", err),
		}
	}

	if err := h.blockerFacade.DeleteUserBlockTx(context.Background(), userUUID); err != nil {
		return h.errorResult(
			err,
			"❌ Ошибка",
			"❌ Не удалось удалить блокировку",
		)
	}

	return CallbackResult{
		ActionText:   "✅ Отклонено",
		ResponseText: "Жалоба успешно отклонена",
	}
}

func (h *BotHandler) errorResult(err error, actionText, responseText string) CallbackResult {
	h.logger.Error(errors.Errorf("request error: error: %s, actionText: %s, responseText: %s", err.Error(), actionText, responseText))
	return CallbackResult{
		Error:        err,
		ActionText:   actionText,
		ResponseText: responseText,
	}
}

func (h *BotHandler) cleanupOnError(ctx context.Context, userUUID uuid.UUID, err error) {
	if err != nil {
		if cleanupErr := h.blockerFacade.DeleteUserBlockTx(ctx, userUUID); cleanupErr != nil {
			h.logger.Error(errors.Errorf("error cleaning user block data: original_error: %s, cleanup_error: %s", err.Error(), cleanupErr.Error()))
		}
	}
}
