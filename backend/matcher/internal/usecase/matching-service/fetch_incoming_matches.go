package matching_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) FetchIncomingMatches(ctx context.Context, in *dto.FetchIncomingMatchesV1Request) (*dto.FetchIncomingMatchesV1Response, error) {
	matches, err := c.matchingStorageFacade.FetchIncomingMatchesTx(ctx, in.ChosenUserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return &dto.FetchIncomingMatchesV1Response{Matches: matches}, nil
}
