package idempotency

import (
	"context"
	"errors"
)

func NewKey(s string) (Key, error) {
	if s == "" {
		return "", errors.New("idempotency key is empty")
	}

	return Key(s), nil
}

type Key string

type ctxKey struct{}

func ContextWithKey(ctx context.Context, key Key) context.Context {
	return context.WithValue(ctx, ctxKey{}, key)
}
func KeyFromContext(ctx context.Context) (Key, bool) {
	tx, ok := ctx.Value(ctxKey{}).(Key)
	return tx, ok
}
