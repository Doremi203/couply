package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewUnaryTokenInterceptor(
	provider Provider,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		token, err := fromGRPCCtx(ctx, provider)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "valid user token is not provided")
		}
		return handler(contextWithToken(ctx, token), req)
	}
}

func fromGRPCCtx(
	ctx context.Context,
	provider Provider,
) (Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return Token{}, errors.Error("failed to extract grpc metadata")
	}

	tokens := md.Get("user-token")
	if len(tokens) == 0 {
		return Token{}, errors.Error("failed to find user-token in grpc metadata")
	}

	token, err := provider.Parse(tokens[0])
	if err != nil {
		return Token{}, errors.WrapFail(err, "parse user-token")
	}

	return token, nil
}
