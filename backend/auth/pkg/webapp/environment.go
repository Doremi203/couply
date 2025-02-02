package webapp

import (
	"strings"
)

type Environment string

const (
	DevEnvironment     Environment = "dev"
	ProdEnvironment    Environment = "prod"
	TestingEnvironment Environment = "testing"
)

func parseEnvironment(s string) Environment {
	switch env := Environment(strings.ToLower(s)); env {
	case DevEnvironment, ProdEnvironment, TestingEnvironment:
		return env
	default:
		return ProdEnvironment
	}
}

func (e Environment) String() string {
	return string(e)
}
