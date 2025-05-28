package worker

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/common/libs/sqs"
	"github.com/Doremi203/couply/backend/matcher/gen/api/messages"
	"github.com/Doremi203/couply/backend/notificator/internal/domain/push"
	"github.com/Doremi203/couply/backend/notificator/internal/usecase"
	"github.com/google/uuid"
)

func NewMatcherEventProcessor(
	logger log.Logger,
	sqsClient sqs.ClientReader[*messages.MatcherEvent],
	senderUseCase usecase.PushSender,
	subscriptionUseCase usecase.PushSubscription,
) *MatcherEventProcessor {
	return &MatcherEventProcessor{
		logger:              logger,
		sqsClient:           sqsClient,
		senderUseCase:       senderUseCase,
		subscriptionUseCase: subscriptionUseCase,
	}
}

type MatcherEventProcessor struct {
	sqsClient           sqs.ClientReader[*messages.MatcherEvent]
	senderUseCase       usecase.PushSender
	subscriptionUseCase usecase.PushSubscription
	logger              log.Logger
}

func (p *MatcherEventProcessor) ProcessMessages(ctx context.Context) error {
	for {
		time.Sleep(time.Second * 5)
		batch, err := p.sqsClient.ReadMessages(ctx, p.logger, 5)
		if err != nil {
			p.logger.Error(errors.WrapFail(err, "read messages batch from sqs"))
			continue
		}

		for _, msg := range batch {
			err := p.processMessage(ctx, msg)
			if err != nil {
				p.logger.Error(errors.WrapFailf(
					err,
					"process message %v",
					errors.Token("message_id", msg.ID),
				))
			}
		}
	}
}

func (p *MatcherEventProcessor) processMessage(
	ctx context.Context,
	msg sqs.Message[*messages.MatcherEvent],
) error {
	defer func() {
		if err := p.sqsClient.DeleteMessage(msg); err != nil {
			p.logger.Error(errors.WrapFailf(
				err,
				"delete message %v",
				errors.Token("message_id", msg.ID),
			))
		}
	}()

	var text string

	switch msg.Data.GetType() {
	case messages.MatcherEvent_MATCH:
		text = "У вас мэтч! Нажмите, чтобы начать общение."

	case messages.MatcherEvent_LIKE:
		text = "Вас кто-то лайкнул! Нажмите, чтобы посмотреть."

	default:
		return errors.Errorf(
			"unknown matcher %v",
			errors.Token("event_type", msg.Data.GetType()),
		)
	}

	idStr, err := uuid.Parse(msg.Data.GetReceiverId())
	if err != nil {
		return errors.WrapFailf(
			err,
			"parse receiver id %v",
			errors.Token("receiver_id", msg.Data.GetReceiverId()),
		)
	}

	pushRecipientID := push.RecipientID(idStr)

	pushRecepient, err := p.subscriptionUseCase.GetRecipient(ctx, pushRecipientID)
	if err != nil {
		return errors.WrapFailf(
			err,
			"get push recipient with id %v",
			errors.Token("id", pushRecipientID),
		)
	}
	if len(pushRecepient.Subscriptions) == 0 {
		p.logger.Infof(
			"no subscriptions for recipient %v, skipping push",
			errors.Token("id", pushRecipientID),
		)
		return nil
	}

	err = p.senderUseCase.Send(ctx, pushRecepient, push.Push{
		Title: "Couply",
		Body:  text,
		Icon:  "/logo.png",
		Url:   "/likes",
	})
	if err != nil {
		return errors.WrapFailf(
			err,
			"send push to recipient with id %v",
			errors.Token("id", pushRecipientID),
		)
	}

	return nil
}
