package errors

import "fmt"

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
	return fmt.Sprintf("[%s]", t.name)
}
