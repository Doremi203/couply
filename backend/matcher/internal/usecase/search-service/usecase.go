package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type searchStorageFacade interface {
	CreateFilterTx(ctx context.Context, newFilter *search.Filter) (*search.Filter, error)
	UpdateFilterTx(ctx context.Context, filter *search.Filter) (*search.Filter, error)
	GetFilterTx(ctx context.Context, userID int64) (*search.Filter, error)
}

type UseCase struct {
	searchStorageFacade searchStorageFacade
}

func NewUseCase(searchStorageFacade searchStorageFacade) *UseCase {
	return &UseCase{searchStorageFacade: searchStorageFacade}
}
