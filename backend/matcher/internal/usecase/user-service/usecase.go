package user_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type userStorageFacade interface {
	userStorageSetterFacade
	userStorageGetterFacade
}

type userStorageSetterFacade interface {
	CreateUserTx(ctx context.Context, user *user.User) error
	UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error)
	DeleteUserTx(ctx context.Context, userID uuid.UUID) error
	UpdatePhotosUploadedAtTx(ctx context.Context, orderNumbers []int32, userID uuid.UUID) error
}

type userStorageGetterFacade interface {
	GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error)
	GetUsersTx(ctx context.Context, userIDs []uuid.UUID) ([]*user.User, error)
}

type UseCase struct {
	photoURLGenerator user.PhotoURLGenerator
	userStorageFacade userStorageFacade
}

func NewUseCase(
	photoURLGenerator user.PhotoURLGenerator,
	userStorageFacade userStorageFacade,
) *UseCase {
	return &UseCase{
		photoURLGenerator: photoURLGenerator,
		userStorageFacade: userStorageFacade,
	}
}
