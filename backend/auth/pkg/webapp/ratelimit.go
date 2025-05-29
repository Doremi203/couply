package webapp

import (
	"context"
	"net"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type rateLimiter interface {
	Add(ctx context.Context, key string) (bool, error)
}

func newUnaryRateLimiterInterceptor(l rateLimiter) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if p, ok := peer.FromContext(ctx); ok {
			if ip, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
				key := info.FullMethod + ":" + ip
				allowed, err := l.Add(ctx, key)
				if err != nil {
					return nil, errors.WrapFail(err, "check rate limit")
				}
				if !allowed {
					return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
				}
			}
		}

		return handler(ctx, req)
	}
}
