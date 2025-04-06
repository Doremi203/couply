package user_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/domain/user"
)

type userStorageFacade interface {
	CreateUserTx(ctx context.Context, user *user.User) (*user.User, error)
	UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error)
	DeleteUserTx(ctx context.Context, id int64) error
	GetUserTx(ctx context.Context, userID int64) (*user.User, error)
}

type UseCase struct {
	userStorageFacade userStorageFacade
}

func NewUseCase(userStorageFacade userStorageFacade) *UseCase {
	return &UseCase{userStorageFacade: userStorageFacade}
}
