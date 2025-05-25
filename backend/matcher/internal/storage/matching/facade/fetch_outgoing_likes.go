package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) FetchOutgoingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error) {
	likes, err := f.storage.FetchLikes(ctx, postgres.FetchLikesOptions{
		SenderUserID: userID,
		Outgoing:     true,
		Limit:        limit,
		Offset:       offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage.FetchLikes")
	}

	return likes, nil
}
