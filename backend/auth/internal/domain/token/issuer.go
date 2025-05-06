package token

import "github.com/Doremi203/couply/backend/auth/internal/domain/user"

type Issuer interface {
	Issue(user user.User) (Token, error)
}
