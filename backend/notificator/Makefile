### Common targets for all services
DB = postgres://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

export GOOSE_DRIVER=postgres
MIGRATE_CREATE = goose create $(name) sql -s
export TESTCONTAINERS_RYUK_DISABLED=true
export BASE_MIGRATIONS_PATH=$(PWD)

check-coverage:
	-go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-go-test-coverage --config=.local.testcoverage.yml
	@rm cover.out

local-build:
	cd cmd/app && go build -ldflags="-s -w"

lint:
	golangci-lint run

generate:
	go generate ./...

migrate/up:
	goose -v -dir "migrations" "$(DB)notificator$(DB_PARAMS)" up

migrate/down:
	goose -v -dir "migrations" "$(DB)notificator$(DB_PARAMS)" reset

migrate/down-one:
	goose -v -dir "migrations" "$(DB)notificator$(DB_PARAMS)" down

migrate/create:
	cd migrations && $(MIGRATE_CREATE)