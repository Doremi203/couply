package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/Doremi203/couply/backend/common/libs/sqs"
	"github.com/Doremi203/couply/backend/matcher/gen/api/messages"

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
	sqs     sqs.ClientWriter[*messages.MatcherEvent]
}

func newLikeProcessor(storage matchingStorageFacade, client sqs.ClientWriter[*messages.MatcherEvent]) likeProcessor {
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
	revertedLike, err := p.storage.GetWaitingLikeTx(ctx, targetUserID, userID)
	if err != nil {
		if !errors.Is(err, matching.ErrLikeNotFound) {
			return nil, errors.Wrap(err, "storage.GetWaitingLikeTx")
		}
	}

	if revertedLike != nil {
		return p.handleMutualLike(ctx, userID, targetUserID, revertedLike.Message)
	}

	return p.handleNewLike(ctx, userID, targetUserID, message)
}

func (p *likeProcess) handleNewLike(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*dto.LikeUserV1Response, error) {
	like := matching.NewLike(userID, targetUserID, message, matching.StatusWaiting)

	if err := p.storage.LikeUserTx(ctx, like); err != nil {
		return nil, errors.Wrap(err, "storage.LikeUserTx")
	}

	if err := p.sqs.SendMessage(ctx, &messages.MatcherEvent{
		Type:       messages.MatcherEvent_LIKE,
		ReceiverId: targetUserID.String(),
		Like: &messages.Like{
			MsgText: message,
		},
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

	if err = p.sqs.SendMessage(ctx, &messages.MatcherEvent{
		Type:       messages.MatcherEvent_MATCH,
		ReceiverId: targetUserID.String(),
		Match:      &messages.Match{},
	}); err != nil {
		return nil, errors.Wrap(err, "sqs.SendMessageToMatchingQueue")
	}

	return &dto.LikeUserV1Response{
		IsMatch: true,
		Match:   newMatch,
	}, nil
}
