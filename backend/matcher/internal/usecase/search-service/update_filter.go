package search_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) UpdateFilter(ctx context.Context, in *dto.UpdateFilterV1Request) (*dto.UpdateFilterV1Response, error) {
	filter := dto.UpdateFilterRequestToFilter(in)

	updatedFilter, err := c.searchStorageFacade.UpdateFilterTx(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateFilterV1Response{Filter: updatedFilter}, nil
}
