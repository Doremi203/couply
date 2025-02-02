package app

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type logFormat string

const (
	consoleLogFormat = "console"
	jsonLogFormat    = "json"
)

func parseLogFormat(s string) (logFormat, error) {
	switch f := logFormat(strings.ToLower(s)); f {
	case consoleLogFormat, jsonLogFormat:
		return f, nil
	default:
		return "", fmt.Errorf("invalid log format: %s", s)
	}
}

func newLogger(format logFormat) *slog.Logger {
	var h slog.Handler

	switch format {
	case consoleLogFormat:
		h = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)
	case jsonLogFormat:
		h = slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelError,
			},
		)
	}

	return slog.New(h)
}
