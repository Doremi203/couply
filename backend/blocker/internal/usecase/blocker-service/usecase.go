//go:generate mockgen -source=usecase.go -destination=../../mocks/usecase/blocker/facade_mock.go -typed

package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/google/uuid"
)

type userClient interface {
	GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error)
	UpdateUserByIDV1(ctx context.Context, user *userservicegrpc.User) error
}

type botClient interface {
	SendReportMessage(user *userservicegrpc.User, reasons []blocker.ReportReason, message string, blockID uuid.UUID) error
	StartCallbackHandler()
}

type blockerStorageFacade interface {
	ReportUserTx(ctx context.Context, block *blocker.UserBlock) error
	DeleteUserBlockTx(ctx context.Context, blockID uuid.UUID) error
	GetBlockInfoTx(ctx context.Context, userID uuid.UUID) (*blocker.UserBlock, error)
}

type UseCase struct {
	userServiceClient    userClient
	bot                  botClient
	blockerStorageFacade blockerStorageFacade
	logger               log.Logger
}

func NewUseCase(userServiceClient userClient, bot botClient, blockerStorageFacade blockerStorageFacade, logger log.Logger) *UseCase {
	return &UseCase{
		userServiceClient:    userServiceClient,
		bot:                  bot,
		blockerStorageFacade: blockerStorageFacade,
		logger:               logger,
	}
}
