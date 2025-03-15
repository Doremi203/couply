### Common targets for all services
DB = postgres://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

MIGRATE_CREATE = migrate create -ext sql -dir

deps/protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest

deps/linters:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.7

common/pull/proto:
	buf export buf.build/googleapis/googleapis --output common/proto/

generate: auth/generate

services/deploy:
	docker-compose up

### Auth service targets
auth/lint:
	cd auth && golangci-lint run

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
