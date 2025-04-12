### Common targets for all services
DB = postgresql://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

MIGRATE_CREATE = migrate create -ext sql -dir

deps/go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

deps/buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

deps/mockgen:
	go install go.uber.org/mock/mockgen@latest

deps/protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest

deps/linters:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.7

common/pull/proto: common/pull/proto/googleapis

common/pull/proto/googleapis:
	$(call CLONE_PROTO_REPO, https://github.com/googleapis/googleapis,google/api/http.proto google/api/annotations.proto)

define CLONE_PROTO_REPO
	@echo "Клонируем $(1)"
	git clone --filter=blob:none --no-checkout $(1) common/proto
	cd common/proto && \
		git sparse-checkout init --no-cone && \
		git sparse-checkout set $(2) && \
		git checkout
	rm -rf common/proto/.git
endef

generate: auth/generate

services/deploy:
	docker-compose up

### Auth service targets
auth/check-coverage:
	-cd auth && MIGRATIONS_PATH=$(PWD)/auth/migrations go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd auth && go-test-coverage --config=.local.testcoverage.yml
	@rm auth/cover.out

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
	migrate -source $(AUTH_MIGRATIONS) -database "$(DB)auth$(DB_PARAMS)" up

auth/migrate/down:
	migrate -source $(AUTH_MIGRATIONS) -database "$(DB)auth$(DB_PARAMS)" down -all

auth/migrate/create:
	$(MIGRATE_CREATE) $(AUTH_MIGRATIONS) $(name)
