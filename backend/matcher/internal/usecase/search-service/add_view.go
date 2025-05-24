package search_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
)

func (c *UseCase) AddView(ctx context.Context, in *dto.AddViewV1Request) (*dto.AddViewV1Response, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = c.searchStorageFacade.AddViewTx(ctx, userID, in.ViewedID)
	if err != nil {
		return nil, err
	}

	return &dto.AddViewV1Response{}, nil
}
