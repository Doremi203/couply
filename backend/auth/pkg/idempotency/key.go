package idempotency

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewKey(s string) (Key, error) {
	_, err := uuid.Parse(s)
	if err != nil {
		return "", errors.Wrap(err, "idempotency key should be a valid UUID")
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
