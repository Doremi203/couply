package phoneconfirm

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

type HashedCode []byte

func NewCodeValue(val string) (CodeValue, error) {
	if len(val) == 0 {
		return "", errors.Error("code value should not be empty")
	}

	return CodeValue(val), nil
}

type CodeValue string

type Code struct {
	value     CodeValue
	ExpiresIn time.Duration
}

func (c Code) Value() CodeValue {
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
		value:     CodeValue(sb.String()),
		ExpiresIn: g.cfg.ExpirationTime,
	}, nil
}
