package user_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type userStorageFacade interface {
	CreateUserTx(ctx context.Context, user *user.User) (*user.User, error)
	UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error)
	DeleteUserTx(ctx context.Context, userID uuid.UUID) error
	GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error)
}

type UseCase struct {
	userStorageFacade userStorageFacade
}

func NewUseCase(userStorageFacade userStorageFacade) *UseCase {
	return &UseCase{userStorageFacade: userStorageFacade}
}
