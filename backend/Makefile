### Common targets for all services
DB = postgres://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

export GOOSE_DRIVER=postgres
MIGRATE_CREATE = goose create $(name) sql -s
export TESTCONTAINERS_RYUK_DISABLED=true
export BASE_MIGRATIONS_PATH=$(PWD)

secrets/yc-token:
	@echo "YC_TOKEN=$$(yc iam create-token)" > secrets.env

deps: deps/migrate deps/go-test-coverage deps/buf deps/mockgen deps/protoc deps/linters

deps/migrate:
	go install github.com/pressly/goose/v3/cmd/goose@latest

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
	brew install golangci-lint

common/pull/proto: common/pull/proto/googleapis

common/pull/proto/googleapis:
	$(call CLONE_PROTO_REPO, https://github.com/googleapis/googleapis,google/api/http.proto google/api/annotations.proto google/api/field_behavior.proto)

common/pull/proto/protoc-gen-openapiv2:
	$(call CLONE_PROTO_REPO, https://github.com/grpc-ecosystem/grpc-gateway,protoc-gen-openapiv2/options/annotations.proto protoc-gen-openapiv2/options/openapiv2.proto)

common/pull/proto/validate:
	$(call CLONE_PROTO_REPO, https://github.com/bufbuild/protoc-gen-validate,validate/validate.proto)

define CLONE_PROTO_REPO
	@echo "Клонируем $(1)"
	@rm -rf tmp
	@git clone --filter=blob:none --no-checkout $(1) tmp
	@cd tmp && \
		git sparse-checkout init --no-cone && \
		git sparse-checkout set $(2) && \
		git checkout
	@mkdir -p common/proto
	@cd tmp && \
		rsync -aR $(2) $(CURDIR)/common/proto/
	@rm -rf tmp
endef

generate: auth/generate matcher/generate blocker/generate notificator/generate

services/deploy:
	docker compose up

lint:
	golangci-lint run

### Auth service targets
auth/check-coverage:
	-cd auth && go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd auth && go-test-coverage --config=.local.testcoverage.yml
	@rm auth/cover.out

auth/local-build:
	cd auth/cmd/app && go build -ldflags="-s -w" -o /dev/null

auth/lint:
	cd auth && golangci-lint run

auth/run:
	docker compose up -d auth

auth/start: auth/generate
	docker compose up --build --force-recreate -d auth

auth/generate:
	cd auth && go generate ./...

AUTH_MIGRATIONS = auth/migrations

auth/migrate/up:
	goose -v -dir "$(AUTH_MIGRATIONS)" "$(DB)auth$(DB_PARAMS)" up

auth/migrate/down:
	goose -v -dir "$(AUTH_MIGRATIONS)" "$(DB)auth$(DB_PARAMS)" reset

auth/migrate/down-one:
	goose -v -dir "$(AUTH_MIGRATIONS)" "$(DB)auth$(DB_PARAMS)" down

auth/migrate/create:
	cd auth/migrations && $(MIGRATE_CREATE)

### Matcher service targets
matcher/check-coverage:
	-cd matcher && go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd matcher && go-test-coverage --config=.local.testcoverage.yml
	@rm matcher/cover.out

matcher/local-build:
	cd matcher/cmd/app && go build -ldflags="-s -w" -o /dev/null

matcher/lint:
	cd matcher && golangci-lint run

matcher/run:
	docker compose up -d matcher

matcher/start: matcher/generate
	docker compose up --build --force-recreate -d matcher

matcher/generate:
	cd matcher && go generate ./...

MATCHER_MIGRATIONS = matcher/migrations

matcher/migrate/up:
	goose -v -dir "$(MATCHER_MIGRATIONS)" "$(DB)matcher$(DB_PARAMS)" up

matcher/migrate/down:
	goose -v -dir "$(MATCHER_MIGRATIONS)" "$(DB)matcher$(DB_PARAMS)" reset

matcher/migrate/down-one:
	goose -v -dir "$(MATCHER_MIGRATIONS)" "$(DB)matcher$(DB_PARAMS)" down

matcher/migrate/create:
	cd matcher/migrations && $(MIGRATE_CREATE)

### Notificator service targets
notificator/check-coverage:
	-cd notificator && go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd notificator && go-test-coverage --config=.local.testcoverage.yml
	@rm notificator/cover.out

notificator/local-build:
	cd notificator/cmd/app && go build -ldflags="-s -w" -o /dev/null

notificator/lint:
	cd notificator && golangci-lint run

notificator/generate:
	cd notificator && go generate ./...

NOTIFICATOR_MIGRATIONS = notificator/migrations

notificator/migrate/up:
	goose -v -dir "$(NOTIFICATOR_MIGRATIONS)" "$(DB)notificator$(DB_PARAMS)" up

notificator/migrate/down:
	goose -v -dir "$(NOTIFICATOR_MIGRATIONS)" "$(DB)notificator$(DB_PARAMS)" reset

notificator/migrate/down-one:
	goose -v -dir "$(NOTIFICATOR_MIGRATIONS)" "$(DB)notificator$(DB_PARAMS)" down

notificator/migrate/create:
	cd $(NOTIFICATOR_MIGRATIONS) && $(MIGRATE_CREATE)


### Blocker service targets
blocker/check-coverage:
	-cd blocker && go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd blocker && go-test-coverage --config=.local.testcoverage.yml
	@rm blocker/cover.out

blocker/local-build:
	cd blocker/cmd/app && go build -ldflags="-s -w" -o /dev/null

blocker/lint:
	cd blocker && golangci-lint run

blocker/run:
	docker compose up -d blocker

blocker/start: blocker/generate
	docker compose up --build --force-recreate -d blocker

blocker/generate:
	cd blocker && go generate ./...

BLOCKER_MIGRATIONS = blocker/migrations

blocker/migrate/up:
	goose -v -dir "$(BLOCKER_MIGRATIONS)" "$(DB)blocker$(DB_PARAMS)" up

blocker/migrate/down:
	goose -v -dir "$(BLOCKER_MIGRATIONS)" "$(DB)blocker$(DB_PARAMS)" reset

blocker/migrate/down-one:
	goose -v -dir "$(BLOCKER_MIGRATIONS)" "$(DB)blocker$(DB_PARAMS)" down

blocker/migrate/create:
	cd $(BLOCKER_MIGRATIONS) && $(MIGRATE_CREATE)