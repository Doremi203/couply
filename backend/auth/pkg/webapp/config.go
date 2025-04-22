package webapp

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
)

type grpcConfig struct {
	Port int
}

type httpConfig struct {
	Port int
}

type loggingConfig struct {
	Level  string
	Format string
}

type swaggerUIConfig struct {
	Path    string
	Enabled bool
}

type Config struct {
	grpc      grpcConfig
	http      httpConfig
	logging   loggingConfig
	swaggerUI swaggerUIConfig

	viperLoader *viper.Viper
	logger      *slog.Logger
}

func (c *Config) ReadSection(name string, cfg any) error {
	err := c.readSection(name, cfg)
	if err != nil {
		return err
	}
	c.logger.Info("loaded custom config", "section", name)

	return nil
}

func (c *Config) readSection(name string, cfg any) error {
	err := c.viperLoader.UnmarshalKey(name, cfg)
	if err != nil {
		return errors.WrapFailf(err, "read section %s", name)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return errors.WrapFailf(err, "read envs for config with %s", name)
	}

	return nil
}

func loadConfig(
	env Environment,
) (Config, error) {
	configsPath := os.Getenv("CONFIGS_PATH")
	fmt.Println("CONFIGS_PATH", configsPath)

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
		return Config{}, fmt.Errorf("no config found in %s", configsPath)
	}

	err := createConfig(v, configs)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		viperLoader: v,
	}

	err = cfg.readSection("grpc", &cfg.grpc)
	if err != nil {
		return Config{}, errors.Wrap(err, "load grpc server config")
	}

	err = cfg.readSection("http", &cfg.http)
	if err != nil {
		return Config{}, errors.Wrap(err, "load http server config")
	}

	err = cfg.readSection("logging", &cfg.logging)
	if err != nil {
		return Config{}, errors.Wrap(err, "load logging config")
	}

	err = cfg.readSection("swagger-ui", &cfg.swaggerUI)
	if err != nil {
		return Config{}, errors.WrapFail(err, "load swagger-ui config")
	}

	return cfg, nil
}

func createConfig(v *viper.Viper, configs []string) error {
	v.SetConfigName(configs[0])
	if err := v.ReadInConfig(); err != nil {
		return errors.WrapFailf(err, "read config: %s", configs[0])
	}
	for i := 1; i < len(configs); i++ {
		v.SetConfigName(configs[i])
		if err := v.MergeInConfig(); err != nil {
			return errors.WrapFailf(err, "merge config: %s", configs[i])
		}
	}

	return nil
}
