package errors

import (
	"fmt"
	"maps"
)

type customError struct {
	wrappedErr error
	msg        message

	tokenNames map[string]bool
}

func (e customError) Error() string {
	if e.wrappedErr == nil {
		return e.msg.String()
	}

	return fmt.Sprintf("%s: %v", e.msg, e.wrappedErr)
}

func (e customError) Unwrap() error {
	return e.wrappedErr
}

func (e customError) Is(err error) bool {
	return e.Error() == err.Error()
}

func Error(msg string) error {
	return Errorf(msg)
}

func Errorf(format string, args ...any) error {
	return create(nil, format, args...)
}

func create(wrapped error, format string, args ...any) customError {
	ret := customError{
		wrappedErr: wrapped,
		msg: message{
			format: format,
			tokens: args,
		},
	}

	wrappedCustom, ok := wrapped.(customError)

	if wrapped == nil || !ok {
		ret.tokenNames = make(map[string]bool, len(args))
		for _, arg := range args {
			customToken, ok := arg.(token)
			if !ok {
				continue
			}
			ret.tokenNames[customToken.name] = true
		}
		return ret
	}

	tokens, tokenNames := remapTokenNames(args, wrappedCustom)

	ret.tokenNames = tokenNames
	ret.msg.tokens = tokens

	return ret
}

func remapTokenNames(tokens []any, err customError) ([]any, map[string]bool) {
	curTokenNames := make(map[string]bool, len(err.tokenNames))
	maps.Copy(curTokenNames, err.tokenNames)

	if len(tokens) == 0 {
		return nil, curTokenNames
	}

	for i := range tokens {
		customToken, ok := tokens[i].(token)
		if !ok {
			continue
		}
		for err.tokenNames[customToken.name] {
			customToken.name = fmt.Sprintf("%s#", customToken.name)
		}
		curTokenNames[customToken.name] = true
		tokens[i] = customToken
	}

	return tokens, curTokenNames
}
