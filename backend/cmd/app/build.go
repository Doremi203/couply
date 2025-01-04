package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Doremi203/Couply/backend/internal/app/mw"
	user_service "github.com/Doremi203/Couply/backend/internal/app/user-service"
	"github.com/Doremi203/Couply/backend/internal/logger"
	"github.com/Doremi203/Couply/backend/internal/storage/facade"
	"github.com/Doremi203/Couply/backend/internal/storage/postgres"
	"github.com/Doremi203/Couply/backend/internal/usecase/user"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

const psqlDSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func Build(ctx context.Context) (*Application, func()) {
	pool, err := connectToDB(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}
	logger.Infof(ctx, "connected to database: %v", psqlDSN)

	txManager := postgres.NewTxManager(pool)

	pgRepository := postgres.NewPgStorage(txManager)
	repositoryFacade := facade.NewStorageFacade(txManager, pgRepository)
	useCase := user.NewUseCase(repositoryFacade)
	impl := user_service.NewImplementation(useCase)

	grpcServer, lis := setupGRPCServer(ctx, impl)
	startGRPCServer(ctx, grpcServer, lis)

	httpServer := setupHTTPServer(ctx, grpcHost)
	startHTTPServer(ctx, httpServer)

	adminHTTPServer := setupAdminServer(ctx, adminHost)
	startAdminServer(ctx, adminHTTPServer)

	shutdownHandlers := []func(){
		func() { shutdownServer(ctx, httpServer, "HTTP server") },
		func() { shutdownServer(ctx, adminHTTPServer, "Admin HTTP server") },
		func() { shutdownGRPCServer(ctx, grpcServer) },
	}

	return &Application{
		grpcServer:       grpcServer,
		httpServer:       httpServer,
		adminHTTPServer:  adminHTTPServer,
		shutdownHandlers: shutdownHandlers,
	}, pool.Close
}

func connectToDB(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, psqlDSN)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func setupGRPCServer(ctx context.Context, impl desc.UserServiceServer) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", grpcHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			mw.Logging,
			mw.Auth,
		),
	)

	reflection.Register(grpcServer)
	desc.RegisterUserServiceServer(grpcServer, impl)

	logger.Infof(ctx, "gRPC server listening on: %v", grpcHost)

	return grpcServer, lis
}

func setupHTTPServer(ctx context.Context, grpcHost string) *http.Server {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, req *http.Request) metadata.MD {
			apiToken := req.Header.Get("x-api-token")
			if apiToken != "" {
				return metadata.Pairs("x-api-token", apiToken)
			}
			return nil
		}),
	)

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handlerWithCors := corsMiddleware(mux)

	err := desc.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcHost, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register order service handler: %v", err)
	}

	logger.Infof(ctx, "http server listening on: %v", httpHost)

	return &http.Server{
		Addr:              httpHost,
		Handler:           handlerWithCors,
		ReadHeaderTimeout: 10 * time.Second,
	}
}

func setupAdminServer(ctx context.Context, adminHost string) *http.Server {
	adminServer := chi.NewRouter()

	adminServer.HandleFunc("/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		b, _ := os.ReadFile("./pkg/order-service/v1/order_service.swagger.json")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(b)
	})

	adminServer.Handle("/*", http.FileServer(http.Dir("./swagger-ui/dist")))

	logger.Infof(ctx, "admin http server listening on: %v", adminHost)

	return &http.Server{
		Addr:              adminHost,
		Handler:           adminServer,
		ReadHeaderTimeout: 10 * time.Second,
	}
}

func shutdownServer(ctx context.Context, server *http.Server, serverName string) {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Infof(ctx, "failed to shutdown %s: %v", serverName, err)
	} else {
		logger.Infof(ctx, "%s gracefully stopped", serverName)
	}
}

func shutdownGRPCServer(ctx context.Context, grpcServer *grpc.Server) {
	grpcServer.GracefulStop()
	logger.Info(ctx, "gRPC server gracefully stopped")
}
