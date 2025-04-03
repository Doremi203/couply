# Main Makefile

# Variables

# Default
LOCAL_BIN = $(CURDIR)/bin
OUT_PATH = $(CURDIR)/pkg

# GO
GO_CMD = go
GO_BUILD = $(GO_CMD) build
GO_CLEAN = $(GO_CMD) clean
GO_TEST = $(GO_CMD) test
GO_TOOL = $(GO_CMD) tool
GO_GET = $(GO_CMD) get
GO_MOD = $(GO_CMD) mod
GO_RUN = $(GO_CMD) run
GO_LINT = golangci-lint run

# Docker
DOCKER_COMPOSE = docker-compose

# Goose
GOOSE = goose

# Squawk
SQUAWK = squawk
SQUAWK_RULES = ban-drop-table,prefer-bigint-over-int,prefer-big-int

# Protolint
PROTOLINT = protolint

# File names and directories
BIN_DIR = bin
APP_NAME = homework
MAIN_FILE = cmd/main.go
COVERAGE_FILE_NAME = coverage.out
COVERAGE_HTML_FILE_NAME = coverage.html
PROTO_DIR = api/user-service/v1

# Postgres
PSQL_DSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# Targets

# Default targets for each app launch
all: deps lint generate build test run clean

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)

deps:
	$(GO_GET) -u ./...
	$(GO_MOD) tidy

run:
	$(GO_RUN) $(MAIN_FILE)

lint:
	$(GO_LINT)
	$(SQUAWK) --exclude=$(SQUAWK_RULES) ./migrations/*.sql
	$(PROTOLINT) lint $(PROTO_DIR)

test:
	$(GO_TEST) -coverprofile=$(COVERAGE_FILE_NAME) ./...
	$(GO_TOOL) cover -html=$(COVERAGE_FILE_NAME) -o $(COVERAGE_HTML_FILE_NAME)

clean:
	$(GO_CLEAN)
	rm -f $(BIN_DIR)/$(APP_NAME)
	rm $(COVERAGE_FILE_NAME)

# Database launch in Docker
compose-up:
	$(DOCKER_COMPOSE) up -d couply_db

compose-down:
	$(DOCKER_COMPOSE) down

compose-stop:
	$(DOCKER_COMPOSE) stop couply_db

compose-start:
	$(DOCKER_COMPOSE) start couply_db

compose-ps:
	$(DOCKER_COMPOSE) ps couply_db

# Migrations launch via Goose
goose-add:
	$(GOOSE) -dir ./migrations postgres $(PSQL_DSN) create rename_me sql

goose-up:
	$(GOOSE) -dir ./migrations postgres $(PSQL_DSN) up

goose-down:
	$(GOOSE) -dir ./migrations postgres $(PSQL_DSN) down

goose-status:
	$(GOOSE) -dir ./migrations postgres $(PSQL_DSN) status

# Proto
bin-deps: .vendor-proto
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@latest

generate:
	mkdir -p ${OUT_PATH}
	protoc --proto_path api --proto_path vendor.protogen \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=${OUT_PATH} --go_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=${OUT_PATH} --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway --grpc-gateway_out ${OUT_PATH} --grpc-gateway_opt paths=source_relative \
		--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 --openapiv2_out=${OUT_PATH} \
		--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:${OUT_PATH}" \
		./api/matching-service/v1/matching_service.proto

.vendor-proto: .vendor-proto/google/protobuf .vendor-proto/google/api .vendor-proto/protoc-gen-openapiv2/options .vendor-proto/validate

.vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/grpc-ecosystem && \
 		cd vendor.protogen/grpc-ecosystem && \
		git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
		git checkout
		mkdir -p vendor.protogen/protoc-gen-openapiv2
		mv vendor.protogen/grpc-ecosystem/protoc-gen-openapiv2/options vendor.protogen/protoc-gen-openapiv2
		rm -rf vendor.protogen/grpc-ecosystem

.vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
		cd vendor.protogen/protobuf &&\
		git sparse-checkout set --no-cone src/google/protobuf &&\
		git checkout
		mkdir -p vendor.protogen/google
		mv vendor.protogen/protobuf/src/google/protobuf vendor.protogen/google
		rm -rf vendor.protogen/protobuf

.vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
 		cd vendor.protogen/googleapis && \
		git sparse-checkout set --no-cone google/api && \
		git checkout
		mkdir -p  vendor.protogen/google
		mv vendor.protogen/googleapis/google/api vendor.protogen/google
		rm -rf vendor.protogen/googleapis

.vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor.protogen/tmp && \
		cd vendor.protogen/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor.protogen/validate
		mv vendor.protogen/tmp/validate vendor.protogen/
		rm -rf vendor.protogen/tmp

.PHONY: all build deps run lint test clean compose-up compose-down compose-stop compose-start compose-ps goose-add goose-up goose-down goose-status bin-deps generate
