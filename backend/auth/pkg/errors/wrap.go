package errors

import (
	"errors"
	"fmt"
)

var New = errors.New
var Is = errors.Is
var As = errors.As

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}
