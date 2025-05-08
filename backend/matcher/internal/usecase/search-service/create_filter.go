package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/utils"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) CreateFilter(ctx context.Context, in *dto.CreateFilterV1Request) (*dto.CreateFilterV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := dto.CreateFilterRequestToFilter(in, userID)

	createdFilter, err := c.searchStorageFacade.CreateFilterTx(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &dto.CreateFilterV1Response{Filter: createdFilter}, nil
}
