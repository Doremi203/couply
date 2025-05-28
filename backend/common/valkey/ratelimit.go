package valkey

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeylimiter"
)

type RateLimiterConfig struct {
	Address   string
	Limit     int
	Window    time.Duration
	KeyPrefix string
	Password  string `secret:"valkey-password"`
}

func NewValkeyRateLimiter(cfg RateLimiterConfig) (*valkeyLimiter, error) {
	cli, err := valkeylimiter.NewRateLimiter(valkeylimiter.RateLimiterOption{
		ClientOption: valkey.ClientOption{
			InitAddress: []string{cfg.Address},
			Password:    cfg.Password,
		},
		KeyPrefix: cfg.KeyPrefix,
		Limit:     cfg.Limit,
		Window:    cfg.Window,
	})
	if err != nil {
		return nil, err
	}
	return &valkeyLimiter{client: cli}, nil
}

type valkeyLimiter struct {
	client valkeylimiter.RateLimiterClient
}

func (v *valkeyLimiter) Add(ctx context.Context, key string) (bool, error) {
	res, err := v.client.Allow(ctx, key)
	if err != nil {
		return false, errors.WrapFail(err, "check rate limit with valkey")
	}
	return res.Allowed, nil
}
