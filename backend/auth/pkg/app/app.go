package app

import (
	"auth/pkg/errors"
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

func Run(startup func(ctx context.Context, appCtx *Context) error, config any) {
	var err error

	flags, err := loadFlags()
	if err != nil {
		panic(errors.Wrap(err, "failed to parse flags"))
	}

	log := newLogger(flags.LogFormat)

	err = loadConfig(flags.ConfigsPath, flags.Env, config)
	if err != nil {
		err = errors.Wrap(err, "failed to load configuration")
		return
	}

	defer func() {
		if err != nil {
			log.Error("app run failed", "error", err)
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	appCtx := newContext(log)
	defer appCtx.close()

	startupTimeout := 5 * time.Second
	startupCtx, timeoutCancel := context.WithTimeoutCause(ctx, startupTimeout, errors.New("startup timeout exceeded"))
	defer timeoutCancel()

	err = startup(startupCtx, appCtx)
	if err != nil {
		return
	}

	var serverWG sync.WaitGroup
	serverWG.Add(2)
	go func() {
		defer serverWG.Done()
		if err := appCtx.grpcServer.listenAndServe(); err != nil {
			log.Error("failed grpc serve", "error", err)
		}
	}()
	go func() {
		defer serverWG.Done()
		if err := appCtx.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("failed http serve", "error", err)
		}
	}()

	<-ctx.Done()
	cancel()

	log.Info("shutting down app")
	log.Info("shutting down servers")

	appCtx.grpcServer.gracefulStop()

	httpShutdownCtx, httpCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer httpCancel()

	if err := appCtx.httpServer.Shutdown(httpShutdownCtx); err != nil {
		log.Error("failed http server shutdown within timeout", "error", err)
	}

	serverWG.Wait()
}

func loadConfig(configsPath string, env Environment, config any) error {
	v := viper.New()

	v.AddConfigPath(configsPath)
	v.SetConfigType("yaml")

	var configs []string
	addIfExists := func(name string) {
		fileName := fmt.Sprintf("%s.yaml", name)
		path := filepath.Join(configsPath, fileName)
		if _, err := os.Stat(path); err == nil {
			configs = append(configs, name)
		}
	}

	addIfExists("base")
	addIfExists(env.String())

	if len(configs) == 0 {
		return fmt.Errorf("no config found in %s", configsPath)
	}

	err := createConfig(v, configs)
	if err != nil {
		return err
	}

	err = v.UnmarshalKey("grpc", &grpcConfig)
	if err != nil {
		return errors.Wrap(err, "failed to decode grpc server config into struct")
	}

	err = v.UnmarshalKey("https", &httpConfig)
	if err != nil {
		return errors.Wrap(err, "failed to decode http server config into struct")
	}

	err = v.UnmarshalKey("app", config)
	if err != nil {
		return errors.Wrap(err, "failed to decode app config into struct")
	}

	return nil
}

func createConfig(v *viper.Viper, configs []string) error {
	v.SetConfigName(configs[0])
	if err := v.ReadInConfig(); err != nil {
		return errors.Wrapf(err, "failed to load config: %s", configs[0])
	}
	for i := 1; i < len(configs); i++ {
		v.SetConfigName(configs[i])
		if err := v.MergeInConfig(); err != nil {
			return errors.Wrapf(err, "failed to merge config: %s", configs[i])
		}
	}

	return nil
}

type Flags struct {
	ConfigsPath string
	Env         Environment
	LogFormat   logFormat
}

func loadFlags() (Flags, error) {
	var (
		flags     Flags
		env       string
		logFormat string
	)

	flag.StringVar(&flags.ConfigsPath, "configs", "", "Path to config files")
	flag.StringVar(&env, "env", "", "Environment name")
	flag.StringVar(&logFormat, "log-format", "json", "Log format")
	flag.Parse()

	if env == "" {
		env = os.Getenv("APP_ENV")
	}
	environment, err := parseEnvironment(env)
	if err != nil {
		return Flags{}, err
	}
	flags.Env = environment

	flags.LogFormat, err = parseLogFormat(logFormat)
	if err != nil {
		return Flags{}, err
	}

	return flags, nil
}
