package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/user"
)

type Token struct {
	userID    user.ID
	userEmail user.Email
}

func (t Token) GetUserID() user.ID {
	return t.userID
}

func (t Token) GetEmail() user.Email {
	return t.userEmail
}

type tokenKey struct{}

func contextWithToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}
func FromContext(ctx context.Context) (Token, bool) {
	tx, ok := ctx.Value(tokenKey{}).(Token)
	return tx, ok
}
