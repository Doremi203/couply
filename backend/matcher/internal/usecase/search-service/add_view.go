package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
)

func (c *UseCase) AddView(ctx context.Context, in *dto.AddViewV1Request) (*dto.AddViewV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	err = c.searchStorageFacade.CreateViewTx(ctx, userID, in.ViewedID)
	if err != nil {
		return nil, errors.Wrap(err, "searchStorageFacade.CreateViewTx")
	}

	return &dto.AddViewV1Response{}, nil
}
