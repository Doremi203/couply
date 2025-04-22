package matching_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) CreateMatch(ctx context.Context, in *dto.CreateMatchV1Request) (*dto.CreateMatchV1Response, error) {
	match := dto.CreateMatchRequestToMatch(in)

	createdMatch, err := c.matchingStorageFacade.CreateMatchTx(ctx, match)
	if err != nil {
		return nil, err
	}

	return &dto.CreateMatchV1Response{Match: createdMatch}, nil
}
