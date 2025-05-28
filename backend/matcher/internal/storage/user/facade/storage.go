package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"
	"github.com/google/uuid"
)

type userServiceStorage interface {
	userStorage
	photoStorage
	interestStorage
}

type userStorage interface {
	CreateUser(ctx context.Context, user *user.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUser(ctx context.Context, opts postgres2.GetUserOptions) (*user.User, error)
	GetUsers(ctx context.Context, opts postgres2.GetUsersOptions) ([]*user.User, error)
	UpdateUser(ctx context.Context, userForUpdate *user.User) error
}

type photoStorage interface {
	CreatePhoto(ctx context.Context, userID uuid.UUID, photo user.Photo) error
	DeletePhotos(ctx context.Context, userID uuid.UUID) error
	GetMultiplePhotos(ctx context.Context, opts postgres2.GetMultiplePhotosOptions) (map[uuid.UUID][]user.Photo, error)
	GetPhotos(ctx context.Context, opts postgres2.GetPhotosOptions) ([]user.Photo, error)
	UpdatePhoto(ctx context.Context, photo user.Photo) error
}

type interestStorage interface {
	CreateInterests(ctx context.Context, userID uuid.UUID, interests *interest.Interest) error
	DeleteInterests(ctx context.Context, userID uuid.UUID) error
	GetInterests(ctx context.Context, opts postgres2.GetInterestsOptions) (*interest.Interest, error)
	GetMultipleInterests(ctx context.Context, opts postgres2.GetMultipleInterestsOptions) (map[uuid.UUID]*interest.Interest, error)
}

type StorageFacadeUser struct {
	txManager storage.TransactionManager
	storage   userServiceStorage
}

func NewStorageFacadeUser(
	txManager storage.TransactionManager,
	pgRepository userServiceStorage,
) *StorageFacadeUser {
	return &StorageFacadeUser{
		txManager: txManager,
		storage:   pgRepository,
	}
}
