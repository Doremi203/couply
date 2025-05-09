package phoneconfirm

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

type HashedCode []byte

type Code struct {
	value     string
	ExpiresIn time.Duration
}

func (c Code) Value() string {
	return c.value
}

type CodeGenerator interface {
	// Generate генерирует криптографически безопасный код.
	Generate() (Code, error)
}

type Config struct {
	Length         int
	ExpirationTime time.Duration
}

func NewDigitCodeGenerator(cfg Config) CodeGenerator {
	return &digitCodeGenerator{
		cfg: cfg,
	}
}

type digitCodeGenerator struct {
	cfg Config
}

func (g *digitCodeGenerator) Generate() (Code, error) {
	var sb strings.Builder
	for i := 0; i < g.cfg.Length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return Code{}, err
		}
		sb.WriteByte(byte('0' + n.Int64()))
	}
	return Code{
		value:     sb.String(),
		ExpiresIn: g.cfg.ExpirationTime,
	}, nil
}
