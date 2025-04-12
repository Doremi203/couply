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

func WrapFail(err error, msg string) error {
	return fmt.Errorf("failed to %s: %w", msg, err)
}

func WrapFailf(err error, format string, args ...interface{}) error {
	return WrapFail(err, fmt.Sprintf(format, args...))
}

func Token(name string, value any) token {
	return token{
		name:  name,
		value: value,
	}
}

type token struct {
	name  string
	value any
}

func (t token) String() string {
	return fmt.Sprintf("{%v: %v}", t.name, t.value)
}
