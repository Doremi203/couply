package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTests(m *testing.M, tester *Tester, migrationsDir string) {
	start := time.Now()
	ctx := context.Background()

	cfg := Config{
		User:     "user",
		Password: "pass",
		Database: "testdb",
		Options:  "sslmode=disable",
	}

	dbPort := "5432/tcp"

	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": cfg.Password,
			"POSTGRES_USER":     cfg.User,
			"POSTGRES_DB":       cfg.Database,
		},
		WaitingFor: wait.ForSQL(nat.Port(dbPort), "postgres", func(host string, port nat.Port) string {
			return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
				cfg.User,
				cfg.Password,
				host,
				port.Port(),
				cfg.Database,
				cfg.Options,
			)
		}).WithStartupTimeout(60 * time.Second).WithPollInterval(2 * time.Second),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("coudn't start db container: %s", err)
	}
	defer func() {
		if err := postgresC.Terminate(ctx); err != nil {
			log.Fatalf("coudn't stop db container: %s", err)
		}
	}()

	host, err := postgresC.Host(ctx)
	if err != nil {
		log.Fatalf("couldn't get db host: %s", err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("couldn't get db port: %s", err)
	}
	cfg.Host = host
	cfg.Port = port.Int()

	sqlDB, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		log.Fatalf("couldn't open db connection: %v", err)
	}
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Printf("couldn't close db connection: %v", err)
		}
	}()

	baseMigrationsPath := os.Getenv("BASE_MIGRATIONS_PATH")
	if baseMigrationsPath == "" {
		log.Fatalf("env variable BASE_MIGRATIONS_PATH not set")
	}

	migrator, err := goose.NewProvider(
		goose.DialectPostgres,
		sqlDB, os.DirFS(filepath.Join(baseMigrationsPath, migrationsDir)),
	)
	if err != nil {
		log.Fatalf("couldn't create db migrations provider: %v", err)
	}

	if _, err = migrator.Up(ctx); err != nil {
		log.Fatalf("couldn't run migrations: %v", err)
	}

	t, err := newTester(ctx, cfg)
	if err != nil {
		log.Fatalf("couldn't create tester: %v", err)
	}

	*tester = t

	fmt.Printf("time to setup db %v", time.Since(start))
	m.Run()
}

func newTester(ctx context.Context, cfg Config) (Tester, error) {
	db, err := NewClient(ctx, cfg)
	if err != nil {
		return Tester{}, err
	}
	return Tester{
		db: db,
	}, nil
}

type Tester struct {
	db *client
}

func (tester Tester) Run(
	t *testing.T,
	name string,
	fixtures []string,
	timeout time.Duration,
	testFunc func(*testing.T, context.Context, Client),
) bool {
	return t.Run(name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		options := pgx.TxOptions{
			IsoLevel:   pgx.ReadCommitted,
			AccessMode: pgx.ReadWrite,
		}
		tx, err := tester.db.BeginTx(ctx, options)
		require.NoError(t, err)

		defer func() {
			rbErr := tx.Rollback(ctx)
			require.NoError(t, rbErr)
		}()

		ctx = ContextWithTx(ctx, tx, options)
		for _, fixture := range fixtures {
			_, err := tester.db.Exec(ctx, fixture)
			require.NoError(t, err)
		}
		testFunc(t, ctx, tester.db)
	})
}
