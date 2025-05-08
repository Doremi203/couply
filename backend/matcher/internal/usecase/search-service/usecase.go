package search_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type searchStorageFacade interface {
	CreateFilterTx(ctx context.Context, newFilter *search.Filter) (*search.Filter, error)
	UpdateFilterTx(ctx context.Context, filter *search.Filter) (*search.Filter, error)
	GetFilterTx(ctx context.Context, userID uuid.UUID) (*search.Filter, error)
	SearchUsersTx(ctx context.Context, userID uuid.UUID, page, limit uint64) ([]*user.User, error)
}

type UseCase struct {
	searchStorageFacade searchStorageFacade
}

func NewUseCase(searchStorageFacade searchStorageFacade) *UseCase {
	return &UseCase{searchStorageFacade: searchStorageFacade}
}
