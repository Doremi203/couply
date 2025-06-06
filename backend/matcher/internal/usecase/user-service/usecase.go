//go:generate mockgen -source=usecase.go -destination=../../mocks/usecase/user/facade_mock.go -typed

package user_service

import (
	"context"

	uuidgen "github.com/Doremi203/couply/backend/auth/pkg/uuid"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type userStorageFacade interface {
	userStorageSetterFacade
	userStorageGetterFacade
}

type userStorageSetterFacade interface {
	CreateUserTx(ctx context.Context, user *user.User) error
	UpdateUserTx(ctx context.Context, user *user.User) error
	DeleteUserTx(ctx context.Context, userID uuid.UUID) error
	UpdatePhotosUploadedAtTx(
		ctx context.Context,
		orderNumbers []int32,
		userID uuid.UUID,
	) error
	UpdateVerificationStatusTx(
		ctx context.Context,
		userID uuid.UUID,
		isVerified bool,
	) error
}

type userStorageGetterFacade interface {
	GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error)
	GetUsersTx(ctx context.Context, userIDs []uuid.UUID) ([]*user.User, error)
}

type UseCase struct {
	photoURLGenerator user.PhotoURLGenerator
	userStorageFacade userStorageFacade
	uuidProvider      uuidgen.Provider
}

func NewUseCase(
	photoURLGenerator user.PhotoURLGenerator,
	userStorageFacade userStorageFacade,
	uuidProvider uuidgen.Provider,
) *UseCase {
	return &UseCase{
		photoURLGenerator: photoURLGenerator,
		userStorageFacade: userStorageFacade,
		uuidProvider:      uuidProvider,
	}
}
