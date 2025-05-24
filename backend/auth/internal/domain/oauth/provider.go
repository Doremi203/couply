package oauth

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

type UserAccount struct {
	Provider       ProviderType
	ProviderUserID ProviderUserID
}

type ProviderType string

const (
	YandexProvider ProviderType = "yandex"
	VKProvider     ProviderType = "vk"
)

type ProviderUserID string

type Provider interface {
	ExchangeCodeForToken(context.Context, Code, State, CodeVerifier, DeviceID) (Token, error)
	FetchUserInfo(context.Context, Token) (UserInfo, error)
}

func NewProviderFactory(
	yandexConfig YandexConfig,
	vkConfig VKConfig,
) ProviderFactory {
	return ProviderFactory{
		yandexConfig: yandexConfig,
		vkConfig:     vkConfig,
	}
}

type ProviderFactory struct {
	yandexConfig YandexConfig
	vkConfig     VKConfig
}

func (f ProviderFactory) New(provider ProviderType) (Provider, error) {
	switch provider {
	case YandexProvider:
		return newYandexProvider(f.yandexConfig), nil
	case VKProvider:
		return newVkProvider(f.vkConfig), nil
	default:
		return nil, errors.Errorf("unknown %v", errors.Token("provider_name", provider))
	}
}
