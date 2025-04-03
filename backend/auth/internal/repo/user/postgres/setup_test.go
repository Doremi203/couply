package userpostgres

import (
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"testing"
)

var tester postgres.Tester

func TestMain(m *testing.M) {
	postgres.SetupTests(m, &tester)
}
