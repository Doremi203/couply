package user

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Details string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user not found with %s", e.Details)
}

var ErrAlreadyExists = errors.New("user already exists")
