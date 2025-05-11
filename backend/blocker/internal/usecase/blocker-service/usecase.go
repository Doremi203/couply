package blocker_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/blocker/internal/client/telegram"
	"github.com/Doremi203/couply/backend/blocker/internal/client/user"
	"github.com/patrickmn/go-cache"
	"time"
)

type UseCase struct {
	userServiceClient *user.Client
	bot               *telegram.BotClient
	logger            log.Logger
	tokenCache        *cache.Cache
}

func NewUseCase(userServiceClient *user.Client, bot *telegram.BotClient, logger log.Logger) *UseCase {
	return &UseCase{
		userServiceClient: userServiceClient,
		bot:               bot,
		logger:            logger,
		tokenCache:        cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *UseCase) GetCache() *cache.Cache {
	return c.tokenCache
}
