package oauth

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

type Provider string

const (
	YandexProvider Provider = "yandex"
)

type ProviderUserID string

type InfoFetcher interface {
	Fetch(context.Context, Token) (UserInfo, error)
}

type InfoFetcherFactory struct{}

func (InfoFetcherFactory) New(provider Provider) (InfoFetcher, error) {
	switch provider {
	case YandexProvider:
		return newYandexFetcher(), nil
	default:
		return nil, errors.Errorf("unknown %v", errors.Token("provider_name", provider))
	}
}
