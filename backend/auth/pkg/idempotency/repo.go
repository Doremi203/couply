package idempotency

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

var ErrAlreadyProcessed = errors.Error("already processed")
var ErrNotBeingProcessed = errors.Error("not being processed")

type Repo interface {
	Create(context.Context, Key) error
	UpdateData(ctx context.Context, key Key, data any) error
	GetData(ctx context.Context, key Key, data any) error
}
