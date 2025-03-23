package argon

import (
	"golang.org/x/crypto/argon2"
)

const (
	timeCost   = 1
	memoryCost = 64 * 1024 // 64 MB
	threads    = 4
	keyLen     = 32 // 32 байта
)

//go:generate mockgen -source=v2provider.go -destination=mock/v2provider_mock.go -typed

type Provider interface {
	Hash(password, salt []byte) []byte
}

type V2Provider struct{}

func (a V2Provider) Hash(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, timeCost, memoryCost, threads, keyLen)
}
