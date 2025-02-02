package main

import (
	"auth/internal/userauth"
	"auth/pkg/app"
	"context"
)

var Config struct {
	Test string
}

func main() {
	app.Run(func(ctx context.Context, appCtx *app.Context) error {
		appCtx.GRPCServer().Register(userauth.NewGRPCService())
		return nil
	}, &Config)
}
