### Common targets for all services
DB = postgres://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

MIGRATE_CREATE = migrate create -ext sql -dir

common/pull/proto:
	buf export buf.build/googleapis/googleapis --output common/proto/

generate: auth/generate

services/deploy:
	docker-compose up

### Auth service targets
auth/build:
	cd auth && go build -o bin/auth ./cmd/auth

auth/run: auth/build
	APP_ENV=dev CONFIGS_PATH=auth/configs ./auth/bin/auth

auth/generate:
	cd auth && go generate ./...

AUTH_MIGRATIONS = file://auth/migrations

auth/migrate/up:
	migrate -source $(AUTH_MIGRATIONS) -database $(DB)auth$(DB_PARAMS) up

auth/migrate/down:
	migrate -source $(AUTH_MIGRATIONS) -database $(DB)auth$(DB_PARAMS) down -all

auth/migrate/create:
	$(MIGRATE_CREATE) $(AUTH_MIGRATIONS) $(name)
