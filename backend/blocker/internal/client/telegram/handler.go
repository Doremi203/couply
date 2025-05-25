package telegram

import (
	"context"
	"strings"
	"time"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/telegram"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/google/uuid"
)

var (
	actionTextBlockError     = "❌ Ошибка блокировки"
	actionTextBlockSuccess   = "✅ Заблокирован"
	actionTextDismissSuccess = "✅ Отклонено"

	responseTextCantGetUserData       = "❌ Не удалось получить данные пользователя"
	responseTextCantUpdateUserStatus  = "❌ Не удалось обновить статус пользователя"
	responseTextCantUpdateBlockStatus = "❌ Не удалось обновить статус блокировки"

	responseTextBlockSuccess = "Пользователь успешно заблокирован"
	responseTextDismissText  = "Жалоба успешно отклонена"
)

type userClient interface {
	GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error)
	UpdateUserByIDV1(ctx context.Context, user *userservicegrpc.User) error
}

type botClient interface {
	RegisterCallbackHandler(prefix string, handler telegram.CallbackHandler)
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

func (h *BotHandler) handleBlockAction(callbackData string) telegram.CallbackResult {
	blockID := strings.TrimPrefix(callbackData, "block_")
	return h.processBlockRequest(blockID)
}

func (h *BotHandler) handleDismissAction(callbackData string) telegram.CallbackResult {
	blockID := strings.TrimPrefix(callbackData, "dismiss_")
	return h.processDismissRequest(blockID)
}

func (h *BotHandler) processBlockRequest(blockID string) (retErr telegram.CallbackResult) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blockUUID, err := uuid.Parse(blockID)
	if err != nil {
		return h.errorResult(
			errors.Wrapf(telegram.ErrInvalidBlockIDFormat, "uuid.Parse: %s", blockID),
			actionTextBlockError,
			responseTextCantUpdateBlockStatus,
		)
	}

	defer h.cleanupOnError(ctx, blockUUID, &retErr)

	blockInfo, err := h.blockerFacade.GetUserBlockByIDTx(ctx, blockUUID)
	if err != nil {
		return h.errorResult(
			errors.Wrapf(err, "blockerFacade.GetUserBlockByIDTx: %s", blockUUID),
			actionTextBlockError,
			responseTextCantUpdateBlockStatus,
		)
	}

	userData, err := h.userClient.GetUserByIDV1(ctx, blockInfo.BlockedID.String())
	if err != nil {
		return h.errorResult(
			errors.Wrapf(err, "userClient.GetUserByIDV1: %s", blockInfo.BlockedID),
			actionTextBlockError,
			responseTextCantGetUserData,
		)
	}

	userData.IsBlocked = true
	if err = h.userClient.UpdateUserByIDV1(ctx, userData); err != nil {
		return h.errorResult(
			errors.Wrapf(err, "userClient.UpdateUserByIDV1: %s", blockInfo.BlockedID),
			actionTextBlockError,
			responseTextCantUpdateUserStatus,
		)
	}

	if err = h.blockerFacade.UpdateUserBlockStatusTx(context.Background(), blockUUID, blocker.BlockStatusAccepted); err != nil {
		return h.errorResult(
			errors.Wrapf(err, "blockerFacade.UpdateUserBlockStatusTx: %s", blockUUID),
			actionTextBlockError,
			responseTextCantUpdateBlockStatus,
		)
	}

	return telegram.CallbackResult{
		ActionText:   actionTextBlockSuccess,
		ResponseText: responseTextBlockSuccess,
	}
}

func (h *BotHandler) processDismissRequest(blockID string) telegram.CallbackResult {
	blockUUID, err := uuid.Parse(blockID)
	if err != nil {
		return h.errorResult(
			errors.Wrapf(telegram.ErrInvalidBlockIDFormat, "uuid.Parse: %s", blockID),
			actionTextBlockError,
			responseTextCantUpdateBlockStatus,
		)
	}

	if err = h.blockerFacade.UpdateUserBlockStatusTx(context.Background(), blockUUID, blocker.BlockStatusDeclined); err != nil {
		return h.errorResult(
			errors.Wrapf(err, "blockerFacade.UpdateUserBlockStatusTx: %s", blockUUID),
			actionTextBlockError,
			responseTextCantUpdateBlockStatus,
		)
	}

	return telegram.CallbackResult{
		ActionText:   actionTextDismissSuccess,
		ResponseText: responseTextDismissText,
	}
}

func (h *BotHandler) errorResult(err error, actionText, responseText string) telegram.CallbackResult {
	h.logger.Error(errors.Wrapf(err, "request error with actionText: %s, responseText: %s", actionText, responseText))
	return telegram.CallbackResult{
		Error:        err,
		ActionText:   actionText,
		ResponseText: responseText,
	}
}

func (h *BotHandler) cleanupOnError(ctx context.Context, blockUUID uuid.UUID, result *telegram.CallbackResult) {
	if result.Error == nil {
		return
	}

	if err := h.blockerFacade.UpdateUserBlockStatusTx(ctx, blockUUID, blocker.BlockStatusDeclined); err != nil {
		h.logger.Error(errors.Wrapf(err, "cleanup error for block id %v", blockUUID))
	}
}
