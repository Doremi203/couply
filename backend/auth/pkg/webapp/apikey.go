package webapp

import (
	"context"
	"slices"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const xAPIKeyHeader = "x-api-key"

type xAPIKeyConfig struct {
	SecretAPIKey string `secret:"x-api-key"`
}

func newUnaryAPIKeyInterceptor(
	secretAPIKey string,
	endpointNames ...string,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		if !slices.Contains(endpointNames, info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.Error("failed to extract grpc metadata")
		}

		values := md.Get(xAPIKeyHeader)
		if len(values) == 0 || values[0] != secretAPIKey {
			return nil, status.Error(codes.Unauthenticated, "valid x-api-key header is not provided")
		}

		return handler(ctx, req)
	}
}
