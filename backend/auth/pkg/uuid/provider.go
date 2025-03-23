package uuid

import "github.com/google/uuid"

//go:generate mockgen -source=provider.go -destination=mock/provider_mock.go -typed

type Provider interface {
	GenerateV7() (uuid.UUID, error)
}

type DefaultProvider struct{}

func (p DefaultProvider) GenerateV7() (uuid.UUID, error) {
	return uuid.NewV7()
}
