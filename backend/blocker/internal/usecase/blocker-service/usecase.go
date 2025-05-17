package blocker_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/blocker/internal/client/telegram"
	"github.com/Doremi203/couply/backend/blocker/internal/client/user"
)

type UseCase struct {
	userServiceClient *user.Client
	bot               *telegram.BotClient
	logger            log.Logger
}

func NewUseCase(userServiceClient *user.Client, bot *telegram.BotClient, logger log.Logger) *UseCase {
	return &UseCase{
		userServiceClient: userServiceClient,
		bot:               bot,
		logger:            logger,
	}
}
