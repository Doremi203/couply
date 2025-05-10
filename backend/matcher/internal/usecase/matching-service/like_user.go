package matching_service

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	matching_storage "github.com/Doremi203/couply/backend/matcher/internal/storage/postgres/matching"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/Doremi203/couply/backend/matcher/utils"
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
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	processor := NewLikeProcessor(c.matchingStorageFacade)
	return processor.ProcessLike(ctx, userID, in.GetTargetUserId(), in.GetMessage())
}

func (p *likeProcessor) ProcessLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	revertedLike, err := p.storage.GetLikeTx(ctx, targetUserID, userID)
	if err != nil {
		if !errors.Is(matching_storage.ErrLikeNotFound, err) {
			return nil, fmt.Errorf("failed to check existing like: %w", err)
		}
	}

	if isMutualLike(revertedLike) {
		return p.handleMutualLike(ctx, userID, targetUserID, revertedLike.GetMessage())
	}

	return p.handleNewLike(ctx, userID, targetUserID, message)
}

func isMutualLike(revertedLike *matching.Like) bool {
	return revertedLike != nil && revertedLike.GetStatus() == matching.StatusWaiting
}

func (p *likeProcessor) handleNewLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	like := matching.NewLike(userID, targetUserID, message, matching.StatusWaiting)

	if _, err := p.storage.LikeUserTx(ctx, like); err != nil {
		return nil, fmt.Errorf("failed to save like: %w", err)
	}

	return &dto.LikeUserV1Response{
		IsMatch: false,
		Match:   nil,
	}, nil
}

func (p *likeProcessor) handleMutualLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	newMatch, err := p.storage.HandleMutualLikeTx(ctx, userID, targetUserID, message)
	if err != nil {
		return nil, fmt.Errorf("failed to handle mutual like: %w", err)
	}

	return &dto.LikeUserV1Response{
		IsMatch: true,
		Match:   newMatch,
	}, nil
}
