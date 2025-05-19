package oauth

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

type ProviderType string

const (
	YandexProvider ProviderType = "yandex"
)

type ProviderUserID string

type Provider interface {
	ExchangeCodeForToken(context.Context, Code, State) (Token, error)
	FetchUserInfo(context.Context, Token) (UserInfo, error)
}

func NewProviderFactory(yandexConfig YandexConfig) ProviderFactory {
	return ProviderFactory{
		yandexConfig: yandexConfig,
	}
}

type ProviderFactory struct {
	yandexConfig YandexConfig
}

func (f ProviderFactory) New(provider ProviderType) (Provider, error) {
	switch provider {
	case YandexProvider:
		return newYandexProvider(f.yandexConfig), nil
	default:
		return nil, errors.Errorf("unknown %v", errors.Token("provider_name", provider))
	}
}
