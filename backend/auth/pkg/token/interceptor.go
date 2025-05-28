package token

import (
	"context"
	"slices"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const InterceptAllMethodsOption = "all-methods"

func NewUnaryTokenInterceptor(
	provider Provider,
	logger log.Logger,
	methods ...string,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		if !slices.Contains(methods, InterceptAllMethodsOption) && !slices.Contains(methods, info.FullMethod) {
			return handler(ctx, req)
		}

		token, err := fromGRPCCtx(ctx, provider)
		if err != nil {
			logger.Error(errors.WrapFail(err, "extract user token from context"))
			return nil, status.Error(codes.Unauthenticated, "valid user token is not provided")
		}
		return handler(ContextWithToken(ctx, token), req)
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
