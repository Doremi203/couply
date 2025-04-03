package mw

import (
	"context"

	user_service "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Auth(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	switch req.(type) {
	case
		user_service.CreateUserV1Request,
		user_service.UpdateUserV1Request:

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "can not parse metadata")
		}

		v := md.Get("x-api-token")
		if len(v) == 0 {
			return nil, status.Error(codes.Unauthenticated, "can not parse x-api-token")
		}
	}

	return handler(ctx, req)
}
