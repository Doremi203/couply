package webapp

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

type grpcService interface {
	RegisterToServer(gRPC *grpc.Server)
}

type grpcGatewayService interface {
	RegisterToGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

// ResourceCloser представляет функцию для закрытия или освобождения ресурса.
// При вызове функция должна выполнить необходимые операции по освобождению и
// вернуть ошибку, если что-то пошло не так.
type ResourceCloser func() error

// BackgroundProcess представляет функцию фонового процесса, которая запускается
// в отдельной горутине. Функция принимает контекст для отслеживания сигнала отмены
// и возвращает ошибку в случае возникновения проблем во время выполнения.
type BackgroundProcess func(ctx context.Context) error

// App представляет основное приложение, включающее в себя настройки,
// логирование, HTTP и gRPC серверы, а также управление фоновыми процессами
// и ресурсами.
type App struct {
	wg sync.WaitGroup

	Log     *slog.Logger
	closers []ResourceCloser

	healthCheckFunc   func() bool
	livenessCheckFunc func() bool

	grpcServer      *grpc.Server
	gatewayOptions  []runtime.ServeMuxOption
	gatewayHandlers []grpcGatewayService
	httpServer      *http.Server

	backgroundProcesses  []BackgroundProcess
	backgroundCtx        context.Context
	backgroundCancelFunc context.CancelCauseFunc

	Env    Environment
	Config Config
}

func initApp() *App {
	envStr := os.Getenv("APP_ENV")
	env := parseEnvironment(envStr)
	fmt.Println("APP_ENV", env)

	cfg, err := loadConfig(env)
	if err != nil {
		panic(errors.WrapFail(err, "load app config"))
	}

	log := newLogger(cfg.logging)
	log.Info("starting service with", "env", env)
	log.Info(
		"loaded app config",
		"grpc_cfg", cfg.grpc,
		"http_cfg", cfg.http,
		"logging_cfg", cfg.logging,
		"swagger-ui", cfg.swaggerUI,
	)
	cfg.logger = log

	backgroundCtx, backgroundCancelFunc := context.WithCancelCause(context.Background())

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	return &App{
		Env:                  env,
		Config:               cfg,
		Log:                  log,
		backgroundCtx:        backgroundCtx,
		backgroundCancelFunc: backgroundCancelFunc,
		grpcServer:           grpcServer,
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", cfg.http.Port),
		},
		healthCheckFunc: func() bool {
			return false
		},
		livenessCheckFunc: func() bool {
			return true
		},
	}
}

// Run запускает приложение. Вызов функции блокируется до получения сигнала от OS о завершении приложения.
//
// Для запуска необходимо выставить переменную окружения APP_ENV, которая определяет окружение приложения. Возможные значения: (dev, tests, testing, prod).
//
// Также необходимо указать переменную окружения CONFIGS_PATH, которая указывает путь к директории с конфигурационными файлами.
func Run(setupFunc func(ctx context.Context, app *App) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	a := initApp()

	defer func() {
		if err := recover(); err != nil {
			a.Log.Error("app crashed with panic", "error", err, "stack", string(debug.Stack()))
			os.Exit(1)
		}
	}()

	err := setupFunc(ctx, a)
	if err != nil {
		a.Log.Error("app setup failed", "error", err)
		os.Exit(1)
	}

	grpcMux := runtime.NewServeMux(a.gatewayOptions...)

	for _, service := range a.gatewayHandlers {
		err := a.registerGatewayHandler(grpcMux, service)
		if err != nil {
			a.Log.Error("app register grpc gateway handler failed", "error", err)
			os.Exit(1)
		}
	}

	a.initGRPCServer()
	a.initHTTPServer(grpcMux)

	a.startBackgroundProcesses()

	a.healthCheckFunc = func() bool {
		return true
	}

	<-ctx.Done()
	cancel()

	a.Log.Info("shutting down app")
	a.Log.Info("shutting down servers")

	a.grpcServer.GracefulStop()

	httpShutdownCtx, httpCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer httpCancel()
	if err := a.httpServer.Shutdown(httpShutdownCtx); err != nil {
		a.Log.Error("failed http server shutdown within timeout", "error", err)
	}

	a.shutDown()
}

func (a *App) AddGatewayOptions(opts ...runtime.ServeMuxOption) {
	a.gatewayOptions = append(a.gatewayOptions, opts...)
}

func (a *App) SetHealthCheck(f func() bool) {
	a.healthCheckFunc = f
}

func (a *App) SetLivenessCheck(f func() bool) {
	a.livenessCheckFunc = f
}

// AddCloser регистрирует функцию закрытия ресурса, которая будет вызвана
// при завершении работы приложения. Все добавленные таким образом функции
// вызываются в обратном порядке (первой вызывается последняя добавленная)
// для корректного освобождения ресурсов.
func (a *App) AddCloser(closer ResourceCloser) {
	a.closers = append(a.closers, closer)
}

// AddBackgroundProcess добавляет фоновый процесс, который будет запущен вместе
// с приложением. Фоновый процесс — это функция, принимающая контекст и возвращающая ошибку.
// Такие процессы используются для обслуживания постоянных задач (например, прослушивания портов, запуск воркеров и т.д.).
func (a *App) AddBackgroundProcess(processor BackgroundProcess) {
	a.backgroundProcesses = append(a.backgroundProcesses, processor)
}

// RegisterGRPCServices регистрирует gRPC-сервис в приложении.
func (a *App) RegisterGRPCServices(services ...grpcService) {
	for _, service := range services {
		a.Log.Info("grpc service registered", "service", a.serviceName(service))
		service.RegisterToServer(a.grpcServer)
	}
}

func (a *App) AddGatewayHandlers(
	services ...grpcGatewayService,
) {
	a.gatewayHandlers = append(a.gatewayHandlers, services...)
}

func (a *App) registerGatewayHandler(
	grpcMux *runtime.ServeMux,
	service grpcGatewayService,
) error {
	err := service.RegisterToGateway(
		a.backgroundCtx,
		grpcMux,
		fmt.Sprintf("localhost:%d", a.Config.grpc.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		return errors.WrapFailf(err, "register gateway handler %v", a.serviceName(service))
	}
	a.Log.Info("gateway handler registered", "service", a.serviceName(service))

	return nil
}

func (a *App) initGRPCServer() {
	a.AddBackgroundProcess(func(ctx context.Context) error {
		a.Log.Info("starting listen on", "port", a.Config.grpc.Port)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Config.grpc.Port))
		if err != nil {
			return errors.WrapFail(err, "start grpc server listener")
		}
		a.Log.Info("starting grpc server on", "port", a.Config.grpc.Port)

		err = a.grpcServer.Serve(listener)
		if err != nil {
			return errors.WrapFail(err, "serve grpc")
		}

		return nil
	})
}

func (a *App) initHTTPServer(grpcMux *runtime.ServeMux) {
	mux := http.NewServeMux()

	if a.Config.swaggerUI.Enabled {
		swaggerUIDir := http.Dir(a.Config.swaggerUI.Path)
		mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(swaggerUIDir)))
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		if !a.healthCheckFunc() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/liveliness", func(w http.ResponseWriter, _ *http.Request) {
		if !a.livenessCheckFunc() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("/", grpcMux)

	a.httpServer.Handler = mux
	a.AddBackgroundProcess(func(ctx context.Context) error {
		a.Log.Info("starting http server on", "address", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return errors.WrapFail(err, "serve http")
		}

		return nil
	})
}

func (a *App) startBackgroundProcesses() {
	for _, processor := range a.backgroundProcesses {
		a.wg.Add(1)
		go func() {
			defer a.wg.Done()
			err := processor(a.backgroundCtx)
			if err != nil && !errors.Is(err, errBackgroundProcessStopped) {
				a.Log.Error("failed to start background", "error", err)
			}
		}()
	}
}

var errBackgroundProcessStopped = errors.New("background process has been stopped")

func (a *App) stopBackgroundProcesses() {
	a.Log.Info("stopping background processes")

	a.backgroundCancelFunc(errBackgroundProcessStopped)
	a.wg.Wait()

	a.Log.Info("background processes stopped")
}

func (a *App) closeResources() {
	a.Log.Info("closing resources")

	for i := len(a.closers) - 1; i >= 0; i-- {
		err := a.closers[i]()
		if err != nil {
			a.Log.Error("failed to close resource", "err", err)
		}
	}

	a.Log.Info("resources closed")
}

func (a *App) shutDown() {
	a.stopBackgroundProcesses()
	a.closeResources()

	a.Log.Info("app shut down")
}

func (a *App) serviceName(service any) string {
	t := reflect.TypeOf(service)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
