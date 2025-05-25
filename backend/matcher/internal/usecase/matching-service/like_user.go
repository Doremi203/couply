package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/token"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/google/uuid"
)

type likeProcessor interface {
	ProcessLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error)
}

type likeProcess struct {
	storage matchingStorageFacade
	sqs     sqsClient
}

func newLikeProcessor(storage matchingStorageFacade, client sqsClient) likeProcessor {
	return &likeProcess{storage: storage, sqs: client}
}

func (c *UseCase) LikeUser(ctx context.Context, in *dto.LikeUserV1Request) (*dto.LikeUserV1Response, error) {
	userID, err := token.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "token.GetUserIDFromContext")
	}

	processor := newLikeProcessor(c.matchingStorageFacade, c.sqsClient)
	return processor.ProcessLike(ctx, userID, in.TargetUserId, in.Message)
}

func (p *likeProcess) ProcessLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
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

func (p *likeProcess) handleNewLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	like := matching.NewLike(userID, targetUserID, message, matching.StatusWaiting)

	if err := p.storage.LikeUserTx(ctx, like); err != nil {
		return nil, errors.Wrap(err, "storage.LikeUserTx")
	}

	if _, err := p.sqs.SendMessageToMatchingQueue(&matching.LikeMessage{
		ReceiverID: targetUserID,
		Message:    message,
	}); err != nil {
		return nil, errors.Wrap(err, "sqs.SendMessageToMatchingQueue")
	}

	return &dto.LikeUserV1Response{
		IsMatch: false,
		Match:   nil,
	}, nil
}

func (p *likeProcess) handleMutualLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	newMatch, err := p.storage.HandleMutualLikeTx(ctx, userID, targetUserID, message)
	if err != nil {
		return nil, errors.Wrap(err, "storage.HandleMutualLikeTx")
	}

	if _, err = p.sqs.SendMessageToMatchingQueue(&matching.LikeMessage{
		ReceiverID: targetUserID,
		Message:    message,
	}); err != nil {
		return nil, errors.Wrap(err, "sqs.SendMessageToMatchingQueue")
	}

	if _, err = p.sqs.SendMessageToMatchingQueue(&matching.MatchMessage{
		FirstUserID:  userID,
		SecondUserID: targetUserID,
	}); err != nil {
		return nil, errors.Wrap(err, "sqs.SendMessageToMatchingQueue")
	}

	return &dto.LikeUserV1Response{
		IsMatch: true,
		Match:   newMatch,
	}, nil
}
