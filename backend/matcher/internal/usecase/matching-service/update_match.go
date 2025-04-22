package matching_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) UpdateMatch(ctx context.Context, in *dto.UpdateMatchV1Request) (*dto.UpdateMatchV1Response, error) {
	match := dto.UpdateMatchRequestToMatch(in)

	updatedMatch, err := c.matchingStorageFacade.UpdateMatchTx(ctx, match)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateMatchV1Response{Match: updatedMatch}, nil
}
