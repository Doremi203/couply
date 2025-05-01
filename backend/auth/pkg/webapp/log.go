package webapp

import (
	"log/slog"
	"os"
	"strings"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
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
