package idempotency

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func FromGRPCCtx(
	ctx context.Context,
) (Key, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	keys := md.Get("idempotency-key")
	if len(keys) == 0 {
		return "", false
	}

	key, err := NewKey(keys[0])
	if err != nil {
		return "", false
	}

	return key, true
}
