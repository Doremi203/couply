package webapp

import (
	"fmt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

type Config struct {
	grpc    grpcConfig
	http    httpConfig
	logging loggingConfig

	viperLoader *viper.Viper
}

func (c *Config) ReadSection(name string, cfg any) error {
	err := c.viperLoader.UnmarshalKey(name, cfg)
	if err != nil {
		return errors.Wrapf(err, "failed to read section %s into %v", name, cfg)
	}

	return nil
}

func loadConfig(
	env Environment,
) (Config, error) {
	configsPath := os.Getenv("CONFIGS_PATH")

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

	err = v.UnmarshalKey("grpc", &cfg.grpc)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to decode grpc server config into struct")
	}

	err = v.UnmarshalKey("http", &cfg.http)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to decode http server config into struct")
	}

	err = v.UnmarshalKey("logging", &cfg.logging)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to decode app config into struct")
	}

	return cfg, nil
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
