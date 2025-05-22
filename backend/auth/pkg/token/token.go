package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/google/uuid"
)

var (
	errTokenNotFound = errors.Error("token not found")
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

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userToken, ok := FromContext(ctx)
	if !ok {
		return uuid.Nil, errors.Wrap(errTokenNotFound, "GetUserIDFromContext")
	}
	return userToken.GetUserID(), nil
}
