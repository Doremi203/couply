package user

import (
	"context"

	"github.com/Doremi203/Couply/backend/internal/domain"
)

type userStorageFacade interface {
	CreateUserTx(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUserTx(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUserTx(ctx context.Context, id int64) error
}

type UseCase struct {
	userStorageFacade userStorageFacade
}

func NewUseCase(userStorageFacade userStorageFacade) *UseCase {
	return &UseCase{userStorageFacade: userStorageFacade}
}
