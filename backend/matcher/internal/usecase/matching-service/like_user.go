package matching_service

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("failed to check existing like: %w", err)
	}

	if isMutualLike(revertedLike) {
		return p.handleMutualLike(ctx, userID, targetUserID, revertedLike.GetMessage())
	}

	return p.handleNewLike(ctx, userID, targetUserID, message)
}

func isMutualLike(revertedLike *matching.Like) bool {
	return revertedLike != nil
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
	updatedLike := matching.NewLike(targetUserID, userID, message, matching.StatusAccepted)
	if _, err := p.storage.UpdateLikeTx(ctx, updatedLike); err != nil {
		return nil, fmt.Errorf("failed to update like: %w", err)
	}

	newLike := matching.NewLike(userID, targetUserID, message, matching.StatusAccepted)
	if _, err := p.storage.LikeUserTx(ctx, newLike); err != nil {
		return nil, fmt.Errorf("failed to save like: %w", err)
	}

	newMatch := matching.NewMatch(userID, targetUserID)
	if _, err := p.storage.AddMatchTx(ctx, newMatch); err != nil {
		return nil, fmt.Errorf("failed to create match: %w", err)
	}

	return &dto.LikeUserV1Response{
		IsMatch: true,
		Match:   newMatch,
	}, nil
}
