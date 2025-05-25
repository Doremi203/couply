package errors

import (
	"errors"
	"fmt"
)

var (
	Is   = errors.Is
	As   = errors.As
	Join = errors.Join
)

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return create(err, format, args...)
}

func WrapFail(err error, msg string) error {
	return WrapFailf(err, msg)
}

func WrapFailf(err error, format string, args ...any) error {
	return Wrapf(err, fmt.Sprintf("failed to %s", format), args...)
}
