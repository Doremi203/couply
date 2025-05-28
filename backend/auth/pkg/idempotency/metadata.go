package idempotency

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc/metadata"
)

func FromGRPCCtx(
	ctx context.Context,
) (Key, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.Error("no idempotency-key header provided")
	}

	keys := md.Get("idempotency-key")
	if len(keys) == 0 {
		return "", errors.Error("no idempotency-key header provided")
	}

	key, err := NewKey(keys[0])
	if err != nil {
		return "", err
	}

	return key, nil
}
