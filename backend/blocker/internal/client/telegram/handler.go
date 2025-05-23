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
	UpdateUserBlockStatusTx(ctx context.Context, blockID uuid.UUID, status blocker.BlockStatus) error
	GetUserBlockByIDTx(ctx context.Context, blockID uuid.UUID) (*blocker.UserBlock, error)
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
	blockID := strings.TrimPrefix(callbackData, "block_")
	return h.processBlockRequest(blockID)
}

func (h *BotHandler) handleDismissAction(callbackData string) CallbackResult {
	blockID := strings.TrimPrefix(callbackData, "dismiss_")
	return h.processDismissRequest(blockID)
}

func (h *BotHandler) processBlockRequest(blockID string) (retErr CallbackResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blockUUID, err := uuid.Parse(blockID)
	if err != nil {
		return CallbackResult{
			Error: fmt.Errorf("invalid block ID format: %w", err),
		}
	}

	defer h.cleanupOnError(ctx, blockUUID, err)

	blockInfo, err := h.blockerFacade.GetUserBlockByIDTx(ctx, blockUUID)
	if err != nil {
		return
	}

	userData, err := h.userClient.GetUserByIDV1(ctx, blockInfo.BlockedID.String())
	if err != nil {
		return h.errorResult(
			err,
			"❌ Ошибка блокировки",
			"❌ Не удалось получить данные пользователя",
		)
	}

	userData.IsBlocked = true
	if err := h.userClient.UpdateUserByIDV1(ctx, userData); err != nil {
		return h.errorResult(
			err,
			"❌ Ошибка блокировки",
			"❌ Не удалось обновить статус пользователя",
		)
	}

	if err := h.blockerFacade.UpdateUserBlockStatusTx(context.Background(), blockUUID, blocker.BlockStatusAccepted); err != nil {
		return h.errorResult(
			err,
			"❌ Ошибка блокировки",
			"❌ Не удалось обновить статус блокировки",
		)
	}

	return CallbackResult{
		ActionText:   "✅ Заблокирован",
		ResponseText: "Пользователь успешно заблокирован",
	}
}

func (h *BotHandler) processDismissRequest(blockID string) CallbackResult {
	blockUUID, err := uuid.Parse(blockID)
	if err != nil {
		return CallbackResult{
			Error: fmt.Errorf("invalid user ID format: %w", err),
		}
	}

	if err := h.blockerFacade.UpdateUserBlockStatusTx(context.Background(), blockUUID, blocker.BlockStatusDeclined); err != nil {
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

func (h *BotHandler) cleanupOnError(ctx context.Context, blockUUID uuid.UUID, err error) {
	if err != nil {
		if cleanupErr := h.blockerFacade.UpdateUserBlockStatusTx(ctx, blockUUID, blocker.BlockStatusDeclined); cleanupErr != nil {
			h.logger.Error(errors.Errorf("error cleaning user block data: original_error: %s, cleanup_error: %s", err.Error(), cleanupErr.Error()))
		}
	}
}
