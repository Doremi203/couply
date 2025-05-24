package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) UpdateFilter(ctx context.Context, in *dto.UpdateFilterV1Request) (*dto.UpdateFilterV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	filter := dto.UpdateFilterRequestToFilter(in, userID)

	if err = c.searchStorageFacade.UpdateFilterTx(ctx, filter); err != nil {
		return nil, errors.Wrap(err, "searchStorageFacade.UpdateFilterTx")
	}

	return &dto.UpdateFilterV1Response{Filter: filter}, nil
}
