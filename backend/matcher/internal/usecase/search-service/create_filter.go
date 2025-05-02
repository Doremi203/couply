package search_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) CreateFilter(ctx context.Context, in *dto.CreateFilterV1Request) (*dto.CreateFilterV1Response, error) {
	filter := dto.CreateFilterRequestToFilter(in)

	createdFilter, err := c.searchStorageFacade.CreateFilterTx(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &dto.CreateFilterV1Response{Filter: createdFilter}, nil
}
