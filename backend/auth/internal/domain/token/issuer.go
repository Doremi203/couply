package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Issuer interface {
	Issue(user.ID) (Token, error)
	IssueRefresh(context.Context, user.ID) (Refresh, error)
	IssuePair(context.Context, user.ID) (Pair, error)
}
