package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/google/uuid"
)

type LikeProcessor interface {
	ProcessLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error)
}

type likeProcessor struct {
	storage matchingStorageFacade
}

func NewLikeProcessor(storage matchingStorageFacade) LikeProcessor {
	return &likeProcessor{storage: storage}
}

func (c *UseCase) LikeUser(ctx context.Context, in *dto.LikeUserV1Request) (*dto.LikeUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	processor := NewLikeProcessor(c.matchingStorageFacade)
	return processor.ProcessLike(ctx, userID, in.TargetUserId, in.Message)
}

func (p *likeProcessor) ProcessLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	revertedLike, err := p.storage.GetLikeTx(ctx, targetUserID, userID)
	if err != nil {
		if !errors.Is(err, matching.ErrLikeNotFound) {
			return nil, errors.Wrap(err, "storage.GetLikeTx")
		}
	}

	if isMutualLike(revertedLike) {
		return p.handleMutualLike(ctx, userID, targetUserID, revertedLike.Message)
	}

	return p.handleNewLike(ctx, userID, targetUserID, message)
}

func isMutualLike(revertedLike *matching.Like) bool {
	return revertedLike != nil && revertedLike.Status == matching.StatusWaiting
}

func (p *likeProcessor) handleNewLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	like := matching.NewLike(userID, targetUserID, message, matching.StatusWaiting)

	if err := p.storage.LikeUserTx(ctx, like); err != nil {
		return nil, errors.Wrap(err, "storage.LikeUserTx")
	}

	return &dto.LikeUserV1Response{
		IsMatch: false,
		Match:   nil,
	}, nil
}

func (p *likeProcessor) handleMutualLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	newMatch, err := p.storage.HandleMutualLikeTx(ctx, userID, targetUserID, message)
	if err != nil {
		return nil, errors.Wrap(err, "storage.HandleMutualLikeTx")
	}

	return &dto.LikeUserV1Response{
		IsMatch: true,
		Match:   newMatch,
	}, nil
}
