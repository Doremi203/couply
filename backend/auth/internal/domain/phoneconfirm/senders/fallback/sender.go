package fallback

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
)

func NewSender(
	senders ...phoneconfirm.CodeSender,
) *sender {
	return &sender{
		senders: senders,
	}
}

type sender struct {
	senders []phoneconfirm.CodeSender
}

func (s *sender) Send(ctx context.Context, logger log.Logger, code phoneconfirm.Code, phone user.Phone) error {
	for _, sender := range s.senders {
		if err := sender.Send(ctx, logger, code, phone); err != nil {
			senderName := fmt.Sprintf("%T", sender)
			logger.Warn(errors.WrapFailf(
				err,
				"send code via %v",
				errors.Token("sender", senderName),
			))
			continue
		}

		return nil
	}

	return errors.Error("all code senders failed to send the code")
}
