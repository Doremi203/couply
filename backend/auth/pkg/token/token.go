package token

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/user"
)

type Provider interface {
	Parse(string) (Token, error)
}

type Token struct {
	userID    user.ID
	userEmail user.Email
}

var ErrInvalidToken = errors.Error("invalid token")

func (t Token) GetUserID() user.ID {
	return t.userID
}

func (t Token) GetEmail() user.Email {
	return t.userEmail
}
