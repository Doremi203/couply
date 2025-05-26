package phoneconfirm

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
)

var ErrUnsupportedPhone = errors.Error("phone number is not supported")

type CodeSender interface {
	Send(context.Context, log.Logger, Code, user.Phone) error
}
