package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"
	"github.com/google/uuid"
)

type blockerServiceStorage interface {
	userBlockStorage
	userBlockReasonStorage
}

type userBlockStorage interface {
	CreateUserBlock(ctx context.Context, block *blocker.UserBlock) error
	DeleteUserBlock(ctx context.Context, opts postgres2.DeleteUserBlockOptions) error
	GetUserBlock(ctx context.Context, opts postgres2.GetUserBlockOptions) (*blocker.UserBlock, error)
	UpdateUserBlock(ctx context.Context, block *blocker.UserBlock) error
}

type userBlockReasonStorage interface {
	CreateUserBlockReason(ctx context.Context, blockID uuid.UUID, reason blocker.ReportReason) error
	GetUserBlockReasons(ctx context.Context, opts postgres2.GetUserBlockReasonsOptions) ([]blocker.ReportReason, error)
}

type StorageFacadeBlocker struct {
	txManager storage.TransactionManager
	storage   blockerServiceStorage
}

func NewStorageFacadeBlocker(
	txManager storage.TransactionManager,
	pgRepository blockerServiceStorage,
) *StorageFacadeBlocker {
	return &StorageFacadeBlocker{
		txManager: txManager,
		storage:   pgRepository,
	}
}
