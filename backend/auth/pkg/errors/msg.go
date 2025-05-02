package errors

import (
	"fmt"
)

type message struct {
	format string
	tokens []any
}

func (e message) String() string {
	if len(e.tokens) == 0 {
		return e.format
	}

	return fmt.Sprintf(e.format, e.tokens...)
}
