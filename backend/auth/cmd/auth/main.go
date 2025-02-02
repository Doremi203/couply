package main

import (
	"auth/internal/userauth"
	"auth/pkg/webapp"
)

func main() {
	app, _ := webapp.New()

	app.RegisterGRPCService(userauth.NewGRPCService())
	app.AddReadinessCheck(func() bool {
		return true
	})

	app.Run()
}
