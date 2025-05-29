package webapp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/pkg/retry/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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

	Log         log.Logger
	httpClient  *resty.Client
	ycSDKClient *ycsdk.SDK

	closers []ResourceCloser

	healthCheckFunc   func() bool
	livenessCheckFunc func() bool

	protectedEndpoints []string
	rateLimiter        rateLimiter

	grpcUnaryInterceptors []grpc.UnaryServerInterceptor
	gatewayOptions        []runtime.ServeMuxOption
	grpcServices          []grpcService
	gatewayHandlers       []grpcGatewayService
	httpServer            *http.Server

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

	logger := newLogger(cfg.logging)
	logger.Infof("starting service with %v", errors.Token("env", env))
	logger.Infof(
		"loaded app config: %v %v %v %v",
		errors.Token("grpc_cfg", cfg.grpc),
		errors.Token("http_cfg", cfg.http),
		errors.Token("logging_cfg", cfg.logging),
		errors.Token("swagger-ui", cfg.swaggerUI),
	)
	cfg.logger = logger

	backgroundCtx, backgroundCancelFunc := context.WithCancelCause(context.Background())

	httpClient := resty.New().
		SetTimeout(5 * time.Second).
		SetRetryCount(5).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)

	sdkClient, err := initYCSdk(httpClient, env)
	if err != nil {
		logger.Error(errors.WrapFail(err, "init yc sdk"))
		os.Exit(1)
	}

	app := &App{
		Env:                  env,
		Config:               cfg,
		Log:                  logger,
		ycSDKClient:          sdkClient,
		httpClient:           httpClient,
		backgroundCtx:        backgroundCtx,
		backgroundCancelFunc: backgroundCancelFunc,
		httpServer: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.http.Port),
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			IdleTimeout:       120 * time.Second,
		},
		healthCheckFunc: func() bool {
			return false
		},
		livenessCheckFunc: func() bool {
			return true
		},
		gatewayOptions: []runtime.ServeMuxOption{
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "idempotency-key", "user-token", xAPIKeyHeader:
					return s, true
				default:
					return runtime.DefaultHeaderMatcher(s)
				}
			}),
			runtime.WithForwardResponseOption(setCookieHeaderMatcher),
			runtime.WithForwardResponseOption(redirectHeaderMatcher),
		},
	}

	err = app.loadSecrets()
	if err != nil {
		logger.Error(errors.WrapFail(err, "load secrets"))
		os.Exit(1)
	}

	return app
}

func initYCSdk(httpClient *resty.Client, env Environment) (*ycsdk.SDK, error) {
	var ycToken string
	if env == TestingEnvironment || env == ProdEnvironment {
		tokenResp := struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
			TokenType   string `json:"token_type"`
		}{}
		resp, err := httpClient.R().
			SetHeader("Metadata-Flavor", "Google").
			SetResult(&tokenResp).
			Get(fmt.Sprintf("http://%s/computeMetadata/v1/instance/service-accounts/default/token", ycsdk.GetMetadataServiceAddr()))
		if err != nil {
			return nil, errors.WrapFail(err, "get yc iam token")
		}
		if !resp.IsSuccess() {
			return nil, errors.Error("got unexpected status code from yc iam token")
		}
		ycToken = tokenResp.AccessToken
	}
	if env == DevEnvironment {
		ycToken = os.Getenv("YC_TOKEN")
	}
	if ycToken == "" {
		return nil, errors.Error("yc token is empty")
	}

	retriesDialOption, err := retry.DefaultRetryDialOption()
	if err != nil {
		return nil, errors.WrapFail(err, "create retry dial option")
	}

	ycSDKClient, err := ycsdk.Build(
		context.Background(),
		ycsdk.Config{Credentials: ycsdk.NewIAMTokenCredentials(ycToken)},
		retriesDialOption,
	)
	if err != nil {
		return nil, errors.WrapFail(err, "build yc sdk")
	}

	return ycSDKClient, nil
}

// Run запускает приложение. Вызов функции блокируется до получения сигнала от OS о завершении приложения.
//
// Для запуска необходимо выставить переменную окружения APP_ENV, которая определяет окружение приложения. Возможные значения: (dev, tests, testing, prod).
//
// Также необходимо указать переменную окружения CONFIGS_PATH, которая указывает путь к директории с конфигурационными файлами.
func Run(setupFunc func(ctx context.Context, app *App) error) {
	a := initApp()

	err := run(a, setupFunc)
	if err != nil {
		a.Log.Error(err)
		a.shutDown()
		os.Exit(1)
	}

	a.shutDown()
}

func run(a *App, setupFunc func(ctx context.Context, app *App) error) (err error) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "app crashed with panic %v", errors.Token("stack", string(debug.Stack())))
		}
	}()

	err = setupFunc(ctx, a)
	if err != nil {
		return errors.Wrap(err, "app setup failed")
	}

	grpcMux := runtime.NewServeMux(a.gatewayOptions...)

	grpcServer, err := a.initGRPCServer(grpcMux)
	a.initHTTPServer(grpcMux)

	a.startBackgroundProcesses()

	a.healthCheckFunc = func() bool {
		return true
	}

	<-ctx.Done()

	a.Log.Infof("shutting down app")
	a.Log.Infof("shutting down servers")

	grpcServer.GracefulStop()

	const httpServerShutdownTimeout = 5 * time.Second
	httpShutdownCtx, httpCancel := context.WithTimeout(context.Background(), httpServerShutdownTimeout)
	defer httpCancel()
	if err := a.httpServer.Shutdown(httpShutdownCtx); err != nil {
		a.Log.Error(errors.WrapFailf(err, "shutdown http server within %v", errors.Token("timeout", httpServerShutdownTimeout)))
	}

	return nil
}

func (a *App) SetRateLimiter(rateLimiter rateLimiter) {
	a.rateLimiter = rateLimiter
}

func (a *App) AddAPIKeyProtectedEndpoints(endpointNames ...string) {
	a.protectedEndpoints = append(a.protectedEndpoints, endpointNames...)
}

func (a *App) HTTPClient() *resty.Client {
	return a.httpClient
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

// AddGRPCUnaryInterceptor добавляет в gRPC-сервер указанные interceptors в переданном порядке.
func (a *App) AddGRPCUnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) {
	a.grpcUnaryInterceptors = append(a.grpcUnaryInterceptors, interceptors...)
}

// RegisterGRPCServices регистрирует gRPC-сервис в приложении.
func (a *App) RegisterGRPCServices(services ...grpcService) {
	a.grpcServices = append(a.grpcServices, services...)
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
		return errors.WrapFailf(
			err,
			"register gateway handler for %v",
			errors.Token("service", a.serviceName(service)),
		)
	}
	a.Log.Infof("gateway handler registered for %v", errors.Token("service", a.serviceName(service)))

	return nil
}

func (a *App) initGRPCServer(grpcMux *runtime.ServeMux) (*grpc.Server, error) {
	var xApiCfg xAPIKeyConfig
	err := a.Config.ReadSection("x-api-key", &xApiCfg)
	if err != nil {
		return nil, errors.WrapFail(err, "read x-api-key config")
	}

	defaultInerceptors := []grpc.UnaryServerInterceptor{
		NewUnaryPanicInterceptor(a.Log),
		NewUnaryInternalErrorLogInterceptor(a.Log),
	}
	if a.rateLimiter != nil {
		defaultInerceptors = append(defaultInerceptors, newUnaryRateLimiterInterceptor(a.rateLimiter, a.Log))
	}
	if xApiCfg.SecretAPIKey != "" {
		defaultInerceptors = append(defaultInerceptors, newUnaryAPIKeyInterceptor(xApiCfg.SecretAPIKey, a.protectedEndpoints...))
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(defaultInerceptors...),
		grpc.ChainUnaryInterceptor(a.grpcUnaryInterceptors...),
	)
	reflection.Register(grpcServer)
	for _, service := range a.grpcServices {
		a.Log.Infof("%v registered", errors.Token("grpc_service", a.serviceName(service)))
		service.RegisterToServer(grpcServer)
	}

	for _, service := range a.gatewayHandlers {
		err := a.registerGatewayHandler(grpcMux, service)
		if err != nil {
			return nil, errors.WrapFail(err, "register grpc gateway handler")
		}
	}

	a.AddBackgroundProcess(func(ctx context.Context) error {
		a.Log.Infof("starting listen on %v", errors.Token("port", a.Config.grpc.Port))
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Config.grpc.Port))
		if err != nil {
			return errors.WrapFail(err, "start grpc server listener")
		}
		a.Log.Infof("starting grpc server on %v", errors.Token("port", a.Config.grpc.Port))

		err = grpcServer.Serve(listener)
		if err != nil {
			return errors.WrapFail(err, "serve grpc")
		}

		return nil
	})

	return grpcServer, nil
}

func (a *App) initHTTPServer(grpcMux *runtime.ServeMux) {
	mux := http.NewServeMux()

	if a.Config.swaggerUI.Enabled {
		swaggerUIDir := http.Dir(a.Config.swaggerUI.Path)
		fileServer := http.FileServer(swaggerUIDir)
		mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age=30, must-revalidate")
			fileServer.ServeHTTP(w, r)
		})))
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

	a.httpServer.Handler = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           600,
	}).Handler(mux)
	a.AddBackgroundProcess(func(ctx context.Context) error {
		a.Log.Infof("starting http server on %v", errors.Token("address", a.httpServer.Addr))
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
				a.Log.Error(errors.WrapFail(err, "start background process"))
			}
		}()
	}
}

var errBackgroundProcessStopped = errors.Error("background process has been stopped")

func (a *App) stopBackgroundProcesses() {
	a.Log.Infof("stopping background processes")

	a.backgroundCancelFunc(errBackgroundProcessStopped)
	a.wg.Wait()

	a.Log.Infof("background processes stopped")
}

func (a *App) closeResources() {
	a.Log.Infof("closing resources")

	for i := len(a.closers) - 1; i >= 0; i-- {
		err := a.closers[i]()
		if err != nil {
			a.Log.Error(errors.WrapFail(err, "close resource"))
		}
	}

	a.Log.Infof("resources closed")
}

func (a *App) shutDown() {
	a.stopBackgroundProcesses()
	a.closeResources()

	a.Log.Infof("app shut down")
}

func (a *App) serviceName(service any) string {
	t := reflect.TypeOf(service)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
