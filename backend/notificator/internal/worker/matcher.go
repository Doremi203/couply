package worker

import (
	"context"

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
		batch, err := p.sqsClient.ReadMessages(ctx, p.logger, 5)
		if err != nil {
			return errors.WrapFail(err, "read messages batch from sqs")
		}
		p.logger.Infof("read %d messages from sqs", len(batch))

		for _, msg := range batch {
			var text string

			switch msg.Data.GetType() {
			case messages.MatcherEvent_MATCH:
				text = "У вас мэтч! Нажмите, чтобы начать общение."

			case messages.MatcherEvent_LIKE:
				text = "Вас кто-то лайкнул! Нажмите, чтобы посмотреть."

			default:
				p.logger.Error(errors.Errorf(
					"unknown matcher %v",
					errors.Token("event_type", msg.Data.GetType()),
				))
				continue
			}

			idStr, err := uuid.Parse(msg.Data.GetReceiverId())
			if err != nil {
				p.logger.Error(errors.WrapFailf(
					err,
					"parse receiver id %v",
					errors.Token("receiver_id", msg.Data.GetReceiverId()),
				))
				continue
			}

			pushRecipientID := push.RecipientID(idStr)

			pushRecepient, err := p.subscriptionUseCase.GetRecipient(ctx, pushRecipientID)
			if err != nil {
				p.logger.Error(errors.WrapFailf(
					err,
					"get push recipient with id %v",
					errors.Token("id", pushRecipientID),
				))
				continue
			}
			if len(pushRecepient.Subscriptions) == 0 {
				p.logger.Infof(
					"no subscriptions for recipient %v, skipping push",
					errors.Token("id", pushRecipientID),
				)
				continue
			}

			err = p.senderUseCase.Send(ctx, pushRecepient, push.Push{
				Title: "Couply",
				Body:  text,
				Icon:  "/logo.png",
			})
			if err != nil {
				p.logger.Error(errors.WrapFailf(
					err,
					"send push to recipient with id %v",
					errors.Token("id", pushRecipientID),
				))
				continue
			}

			if err := p.sqsClient.DeleteMessage(msg); err != nil {
				p.logger.Error(errors.WrapFailf(
					err,
					"delete message %v",
					errors.Token("message_id", msg.ID),
				))
				continue
			}
		}
	}
}
