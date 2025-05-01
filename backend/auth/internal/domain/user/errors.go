package user

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user not found: %v", e.Err)
}

var ErrAlreadyExists = errors.New("user already exists")
