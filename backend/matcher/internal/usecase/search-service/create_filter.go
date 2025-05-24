package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) CreateFilter(ctx context.Context, in *dto.CreateFilterV1Request) (*dto.CreateFilterV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	filter := dto.CreateFilterRequestToFilter(in, userID)

	err = c.searchStorageFacade.CreateFilterTx(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "searchStorageFacade.CreateFilterTx")
	}

	return &dto.CreateFilterV1Response{Filter: filter}, nil
}
