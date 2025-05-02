package errors

import (
	"log/slog"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
)

func Logger(base *slog.Logger) log.Logger {
	return &logger{
		base: base,
	}
}

type logger struct {
	base *slog.Logger
}

func (l *logger) Info(format string, args ...any) {
	err := Errorf(format, args...)
	msg, logArgs := l.msgAndArgs(err)
	l.base.Info(msg, logArgs...)
}

func (l *logger) Warn(err error) {
	msg, args := l.msgAndArgs(err)
	l.base.Warn(msg, args...)
}

func (l *logger) Error(err error) {
	msg, args := l.msgAndArgs(err)
	l.base.Error(msg, args...)
}

func (l *logger) msgAndArgs(err error) (string, []any) {
	var customErr customError
	if !As(err, &customErr) {
		l.base.Error(err.Error())
	}

	return customErr.Error(), l.errorArgs(customErr)
}

func (l *logger) errorArgs(err error) []any {
	var customErr customError
	if !As(err, &customErr) || err == nil {
		return nil
	}
	internalRes := l.errorArgs(customErr.wrappedErr)
	ret := make([]any, 0, len(customErr.msg.tokens)*2)
	for _, t := range customErr.msg.tokens {
		t, ok := t.(token)
		if !ok {
			continue
		}
		ret = append(ret, t.name)
		ret = append(ret, t.value)
	}

	return append(ret, internalRes...)
}
