package app

import (
	"errors"
	"strings"
)

type Environment string

const (
	DevEnvironment     Environment = "dev"
	ProdEnvironment    Environment = "prod"
	TestingEnvironment Environment = "testing"
)

func parseEnvironment(s string) (Environment, error) {
	switch env := Environment(strings.ToLower(s)); env {
	case DevEnvironment, ProdEnvironment, TestingEnvironment:
		return env, nil
	default:
		return "", errors.New("unknown environment")
	}
}

func (e Environment) String() string {
	return string(e)
}
