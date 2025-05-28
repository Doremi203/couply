package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"
	"github.com/google/uuid"
)

type matchingServiceStorage interface {
	likeStorage
	matchStorage
}

type likeStorage interface {
	CreateLike(ctx context.Context, like *matching.Like) error
	FetchLikes(ctx context.Context, opts postgres2.FetchLikesOptions) ([]*matching.Like, error)
	GetLike(ctx context.Context, opts postgres2.GetLikeOptions) (*matching.Like, error)
	UpdateLike(ctx context.Context, like *matching.Like) error
}

type matchStorage interface {
	CreateMatch(ctx context.Context, match *matching.Match) error
	DeleteMatch(ctx context.Context, userID, targetUserID uuid.UUID) error
	FetchMatches(ctx context.Context, opts postgres2.FetchMatchesOptions) ([]*matching.Match, error)
	GetMatch(ctx context.Context, opts postgres2.GetMatchOptions) (*matching.Match, error)
}

type StorageFacadeMatching struct {
	txManager storage.TransactionManager
	storage   matchingServiceStorage
}

func NewStorageFacadeMatching(
	txManager storage.TransactionManager,
	pgRepository matchingServiceStorage,
) *StorageFacadeMatching {
	return &StorageFacadeMatching{
		txManager: txManager,
		storage:   pgRepository,
	}
}
