package user

import (
	"errors"
)

var ErrOAuthAccountNotFound = errors.New("oauth account not found")
var ErrOAuthAccountAlreadyExists = errors.New("oauth account already exists")
var ErrNotFound = errors.New("user not found")

var ErrAlreadyExists = errors.New("user already exists")
