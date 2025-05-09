package timeprovider

import "time"

type Provider interface {
	Now() time.Time
}

type ProviderFunc func() time.Time

func (f ProviderFunc) Now() time.Time {
	return f()
}
