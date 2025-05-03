package search_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) GetFilter(ctx context.Context, in *dto.GetFilterV1Request) (*dto.GetFilterV1Response, error) {
	filter, err := c.searchStorageFacade.GetFilterTx(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.GetFilterV1Response{Filter: filter}, nil
}
