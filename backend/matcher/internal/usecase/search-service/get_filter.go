package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/utils"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) GetFilter(ctx context.Context, in *dto.GetFilterV1Request) (*dto.GetFilterV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	filter, err := c.searchStorageFacade.GetFilterTx(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.GetFilterV1Response{Filter: filter}, nil
}
