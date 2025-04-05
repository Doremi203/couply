package matching_service

import (
	"context"
	dto "github.com/Doremi203/Couply/backend/internal/dto/matching-service"
)

func (c *UseCase) FetchOutgoingMatches(ctx context.Context, in *dto.FetchOutgoingMatchesV1Request) (*dto.FetchOutgoingMatchesV1Response, error) {
	matches, err := c.matchingStorageFacade.FetchOutgoingMatchesTx(ctx, in.MainUserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return &dto.FetchOutgoingMatchesV1Response{Matches: matches}, nil
}
