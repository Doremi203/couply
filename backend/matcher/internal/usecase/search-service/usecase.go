//go:generate mockgen -source=usecase.go -destination=../../mocks/usecase/search/facade_mock.go -typed

package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type searchStorageFacade interface {
	searchStorageSetterFacade
	searchStorageGetterFacade
}

type searchStorageSetterFacade interface {
	CreateFilterTx(ctx context.Context, newFilter *search.Filter) error
	UpdateFilterTx(ctx context.Context, filter *search.Filter) error
	CreateViewTx(ctx context.Context, viewerID, viewedID uuid.UUID) error
}

type searchStorageGetterFacade interface {
	GetFilterTx(ctx context.Context, userID uuid.UUID) (*search.Filter, error)
	SearchUsersTx(ctx context.Context, userID uuid.UUID, offset, limit uint64) ([]*user.User, map[uuid.UUID]float64, error)
}

type UseCase struct {
	searchStorageFacade searchStorageFacade
	photoURLGenerator   user.PhotoURLGenerator
	logger              log.Logger
}

func NewUseCase(
	searchStorageFacade searchStorageFacade,
	photoURLGenerator user.PhotoURLGenerator,
	logger log.Logger,
) *UseCase {
	return &UseCase{
		searchStorageFacade: searchStorageFacade,
		photoURLGenerator:   photoURLGenerator,
		logger:              logger,
	}
}
