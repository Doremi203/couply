package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres3 "github.com/Doremi203/couply/backend/matcher/internal/storage/search/postgres"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"
	"github.com/google/uuid"
)

type userServiceStorage interface {
	GetUser(ctx context.Context, opts postgres2.GetUserOptions) (*user.User, error)
	GetPhotos(ctx context.Context, opts postgres2.GetPhotosOptions) ([]user.Photo, error)
	GetInterests(ctx context.Context, opts postgres2.GetInterestsOptions) (*interest.Interest, error)
}

type searchServiceStorage interface {
	CreateUserView(ctx context.Context, viewerID, viewedID uuid.UUID) error
	SearchUsers(
		ctx context.Context,
		filter *search.Filter,
		interests *interest.Interest,
		curLatitude, curLongitude float64,
		offset, limit uint64,
	) ([]*user.User, map[uuid.UUID]float64, error)
	filterStorage
	filterInterestStorage
}

type filterStorage interface {
	CreateFilter(ctx context.Context, filter *search.Filter) error
	GetFilter(ctx context.Context, opts postgres3.GetFilterOptions) (*search.Filter, error)
	UpdateFilter(ctx context.Context, filter *search.Filter) error
}

type filterInterestStorage interface {
	CreateFilterInterests(ctx context.Context, userID uuid.UUID, filterInterests *interest.Interest) error
	DeleteFilterInterests(ctx context.Context, userID uuid.UUID) error
	GetFilterInterests(ctx context.Context, opts postgres3.GetFilterInterestsOptions) (*interest.Interest, error)
}

type StorageFacadeSearch struct {
	txManager     storage.TransactionManager
	searchStorage searchServiceStorage
	userStorage   userServiceStorage
}

func NewStorageFacadeSearch(
	txManager storage.TransactionManager,
	pgRepositorySearch searchServiceStorage,
	pgRepositoryUser userServiceStorage,
) *StorageFacadeSearch {
	return &StorageFacadeSearch{
		txManager:     txManager,
		searchStorage: pgRepositorySearch,
		userStorage:   pgRepositoryUser,
	}
}
