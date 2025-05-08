package postgres

import "fmt"

type Config struct {
	Host     string
	Port     int
	User     string `env:"DATABASE_USER" secret:"db-user"`
	Password string `env:"DATABASE_PASSWORD" secret:"db-password"`
	Database string
	Options  string
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Options,
	)
}
