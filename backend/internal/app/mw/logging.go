package mw

import (
	"context"

	"github.com/Doremi203/Couply/backend/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		logger.Infof(ctx, "[interceptor.Logging] method: %s; metadata: %v", info.FullMethod, md)
	}

	rewReq, err := protojson.Marshal((req).(proto.Message))
	if err != nil {
		logger.Warnf(ctx, "[interceptor.Logging] failed to marshal request: %v", err)
	}
	logger.Infof(ctx, "[interceptor.Logging] method: %s; request: %s", info.FullMethod, string(rewReq))

	res, err := handler(ctx, req)
	if err != nil {
		logger.Warnf(ctx, "[interceptor.Logging] method: %s; error: %s", info.FullMethod, err.Error())
		return
	}

	respReq, err := protojson.Marshal((res).(proto.Message))
	if err != nil {
		logger.Warnf(ctx, "[interceptor.Logging] failed to marshal response: %v", err)
	}
	logger.Infof(ctx, "[interceptor.Logging] method: %s; response: %s", info.FullMethod, string(respReq))

	return res, nil
}
