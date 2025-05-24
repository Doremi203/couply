package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
)

func (c *UseCase) DislikeUser(ctx context.Context, in *dto.DislikeUserV1Request) (*dto.DislikeUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	// message value doesnt matter
	updatedLike := matching.NewLike(in.TargetUserID, userID, "", matching.StatusDeclined)

	if err = c.matchingStorageFacade.UpdateLikeTx(ctx, updatedLike); err != nil {
		return nil, errors.Wrap(err, "matchingStorageFacade.UpdateLikeTx")
	}

	return &dto.DislikeUserV1Response{}, nil
}
