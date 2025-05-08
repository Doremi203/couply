package webapp

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/Doremi203/couply/backend/auth/pkg/log"

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

type secretsConfig struct {
	Ids map[string]string
}

type Config struct {
	grpc      grpcConfig
	http      httpConfig
	logging   loggingConfig
	swaggerUI swaggerUIConfig
	secrets   secretsConfig

	secretsMap map[string]string

	viperLoader *viper.Viper
	logger      log.Logger
}

func (c *Config) ReadSection(name string, cfg any) error {
	err := c.readSection(name, cfg)
	if err != nil {
		return err
	}
	c.logger.Infof("loaded custom config %v", errors.Token("section", name))

	return nil
}

func (c *Config) readSection(name string, cfg any) error {
	err := c.viperLoader.UnmarshalKey(name, cfg)
	if err != nil {
		return errors.WrapFailf(err, "read section %v", errors.Token("name", name))
	}

	err = c.readFromSecrets(cfg)
	if err != nil {
		return errors.WrapFailf(err, "read secrets for config with %v", errors.Token("name", name))
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return errors.WrapFailf(err, "read envs for config with %v", errors.Token("name", name))
	}

	return nil
}

func (c *Config) readFromSecrets(cfg any) error {
	if len(c.secretsMap) == 0 {
		return nil
	}
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.Errorf("want pointer to struct, got %T", cfg)
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get("secret")
		if tagValue == "" {
			continue
		}

		secret, ok := c.secretsMap[tagValue]
		if !ok {
			continue
		}

		fv := v.Field(i)
		if !fv.CanSet() {
			return errors.Errorf("field %q could not be set", field.Name)
		}
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(secret)
		default:
			return fmt.Errorf("field type %q (%s) not supported", field.Name, fv.Kind())
		}
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
		secretsMap:  make(map[string]string),
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

	err = cfg.readSection("secrets", &cfg.secrets)
	if err != nil {
		return Config{}, errors.WrapFail(err, "load secrets config")
	}

	return cfg, nil
}

func createConfig(v *viper.Viper, configs []string) error {
	v.SetConfigName(configs[0])
	if err := v.ReadInConfig(); err != nil {
		return errors.WrapFailf(err, "read %v", errors.Token("config", configs[0]))
	}
	for i := 1; i < len(configs); i++ {
		v.SetConfigName(configs[i])
		if err := v.MergeInConfig(); err != nil {
			return errors.WrapFailf(err, "merge %v", errors.Token("config", configs[i]))
		}
	}

	return nil
}
