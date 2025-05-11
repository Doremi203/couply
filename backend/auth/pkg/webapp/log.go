package webapp

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type logFormat string

const (
	consoleLogFormat = "console"
	jsonLogFormat    = "json"
)

func parseLogFormat(s string) logFormat {
	switch f := logFormat(strings.ToLower(s)); f {
	case consoleLogFormat, jsonLogFormat:
		return f
	default:
		return jsonLogFormat
	}
}

func newLogger(config loggingConfig) log.Logger {
	format := parseLogFormat(config.Format)
	level := parseLogLevel(config.Level)

	var h slog.Handler
	handlerOptions := &slog.HandlerOptions{
		Level: level,
	}

	switch format {
	case consoleLogFormat:
		h = slog.NewTextHandler(
			os.Stdout,
			handlerOptions,
		)
	case jsonLogFormat:
		h = slog.NewJSONHandler(
			os.Stdout,
			handlerOptions,
		)
	}

	return errors.Logger(slog.New(h))
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelError
	}
}

func NewUnaryInternalErrorLogInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			s, ok := status.FromError(err)
			if !ok && s.Code() == codes.Unknown {
				logger.Error(errors.Wrapf(err, "%v failed", errors.Token("handler", info.FullMethod)))
				return resp, status.Error(codes.Internal, "internal error")
			}
		}

		return resp, err
	}
}
