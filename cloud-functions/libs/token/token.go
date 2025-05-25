package token

import (
	"context"

	"github.com/google/uuid"
)

type Token struct {
	userID uuid.UUID
}

func (t Token) GetUserID() uuid.UUID {
	return t.userID
}

type tokenKey struct{}

func contextWithToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}
func FromContext(ctx context.Context) (Token, bool) {
	tx, ok := ctx.Value(tokenKey{}).(Token)
	return tx, ok
}
