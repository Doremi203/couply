package main

import (
	"context"
	"os/signal"
	"syscall"

	application "github.com/Doremi203/Couply/backend/cmd/app"

	"github.com/Doremi203/Couply/backend/internal/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app, poolClose := application.Build(ctx)
	defer poolClose()

	<-ctx.Done()
	logger.Info(ctx, "gracefully shutting down")

	app.Shutdown()
	logger.Info(ctx, "application gracefully stopped")
}
