package pushpostgres

import (
	"testing"

	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
)

var tester postgres.Tester

func TestMain(m *testing.M) {
	postgres.SetupTests(m, &tester, "notificator")
}
