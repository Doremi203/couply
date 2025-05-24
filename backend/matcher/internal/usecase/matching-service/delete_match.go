package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) DeleteMatch(ctx context.Context, in *dto.DeleteMatchV1Request) (*dto.DeleteMatchV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	err = c.matchingStorageFacade.DeleteMatchTx(ctx, userID, in.TargetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "matchingStorageFacade.DeleteMatchTx")
	}

	return &dto.DeleteMatchV1Response{}, nil
}
