package app

import (
	"net/http"

	"google.golang.org/grpc"
)

type Application struct {
	grpcServer       *grpc.Server
	httpServer       *http.Server
	adminHTTPServer  *http.Server
	shutdownHandlers []func()
}

func (a *Application) Shutdown() {
	for _, handler := range a.shutdownHandlers {
		handler()
	}
}
