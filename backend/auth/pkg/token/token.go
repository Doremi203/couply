package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/google/uuid"
)

var (
	ErrTokenNotFound = errors.Error("token not found")
)

type Token struct {
	UserID uuid.UUID
}

func (t Token) GetUserID() uuid.UUID {
	return t.UserID
}

type tokenKey struct{}

func ContextWithToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}
func FromContext(ctx context.Context) (Token, bool) {
	tx, ok := ctx.Value(tokenKey{}).(Token)
	return tx, ok
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userToken, ok := FromContext(ctx)
	if !ok {
		return uuid.Nil, errors.Wrap(ErrTokenNotFound, "GetUserIDFromContext")
	}
	return userToken.GetUserID(), nil
}
