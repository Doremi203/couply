package idempotency

import "errors"

func NewKey(s string) (Key, error) {
	if s == "" {
		return "", errors.New("idempotency key is empty")
	}

	return Key(s), nil
}

type Key string
