package sms

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
)

var ErrUnsupportedPhoneOperator = errors.Error("phone operator not supported")

type Sender interface {
	Send(ctx context.Context, text, phoneE164 string) error
}

type TestSender struct {
	Logger log.Logger
}

func (s TestSender) Send(ctx context.Context, text, phoneE164 string) error {
	s.Logger.Infof(
		"Sending %v to %v",
		errors.Token("msg", text),
		errors.Token("phone", phoneE164),
	)
	return nil
}
