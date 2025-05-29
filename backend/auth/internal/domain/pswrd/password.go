package pswrd

import (
	"errors"
	"strings"
)

func NewPassword(s string) (Password, error) {
	if len(s) < 6 {
		return "", errors.New("password must be at least 6 characters long")
	}
	if len(s) > 32 {
		return "", errors.New("password must be at most 16 characters long")
	}
	if !strings.ContainsAny(s, "_-!@#?") {
		return "", errors.New("password must contain at least one special character (_-!@#?)")
	}
	if strings.ToLower(s) == s {
		return "", errors.New("password must contain at least one uppercase letter")
	}

	return Password(s), nil
}

type Password string

type HashedPassword []byte
