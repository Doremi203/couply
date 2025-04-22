package matching_service

import (
	"context"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) FetchMatches(ctx context.Context, in *dto.FetchMatchesV1Request) (*dto.FetchMatchesV1Response, error) {
	matches, err := c.matchingStorageFacade.FetchMatchesTx(ctx, in.MainUserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return &dto.FetchMatchesV1Response{Matches: matches}, nil
}
