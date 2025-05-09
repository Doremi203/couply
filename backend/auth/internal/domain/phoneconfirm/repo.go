package phoneconfirm

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Repo interface {
	UpsertRequest(context.Context, Request) error
	GetRequest(context.Context, user.ID, user.Phone) (*Request, error)
}
