package user

import "errors"

var ErrInvalidPassword = errors.New("invalid password")
var ErrNotFound = errors.New("user not found")
var ErrAlreadyExists = errors.New("user already exists")
