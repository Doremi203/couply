package postgres

import "fmt"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
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
