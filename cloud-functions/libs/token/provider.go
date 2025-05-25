package token

type Provider interface {
	Parse(string) (Token, error)
}

type Config struct {
	SecretKey string `env:"SECRET_KEY" secret:"secret-key"`
}
