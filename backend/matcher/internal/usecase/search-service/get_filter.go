package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) GetFilter(ctx context.Context, in *dto.GetFilterV1Request) (*dto.GetFilterV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	filter, err := c.searchStorageFacade.GetFilterTx(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "searchStorageFacade.GetFilterTx")
	}

	return &dto.GetFilterV1Response{Filter: filter}, nil
}
