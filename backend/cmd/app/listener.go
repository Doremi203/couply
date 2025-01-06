package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

const (
	adminHost = ":7003"
	grpcHost  = ":7002"
	httpHost  = ":7001"
)

func startGRPCServer(ctx context.Context, grpcServer *grpc.Server, lis net.Listener) {
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			select {
			case <-ctx.Done():
			default:
				log.Fatalf("failed to serve: %v\n", err)
			}
		}
	}()
}

func startHTTPServer(ctx context.Context, httpServer *http.Server) {
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			select {
			case <-ctx.Done():
			default:
				log.Fatalf("failed to listen and serve order service handler: %v\n", err)
			}
		}
	}()
}

func startAdminServer(ctx context.Context, adminHTTPServer *http.Server) {
	go func() {
		if err := adminHTTPServer.ListenAndServe(); err != nil {
			select {
			case <-ctx.Done():
			default:
				log.Fatalf("failed to listen and serve admin server: %v\n", err)
			}
		}
	}()
}
