### Common targets for all services
DB = postgres://user:pass@localhost:5432/
DB_PARAMS = ?sslmode=disable

MIGRATE_CREATE = migrate create -ext sql -dir
export TESTCONTAINERS_RYUK_DISABLED=true
export BASE_MIGRATIONS_PATH=$(PWD)

deps: deps/migrate deps/go-test-coverage deps/buf deps/mockgen deps/protoc deps/linters

deps/migrate:
	brew install golang-migrate

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

generate: auth/generate matcher/generate

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
	cd auth/cmd/auth && go build -ldflags="-s -w"

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
	migrate -verbose -source $(AUTH_MIGRATIONS) -database "$(DB)auth$(DB_PARAMS)" up

auth/migrate/down:
	migrate -source $(AUTH_MIGRATIONS) -database "$(DB)auth$(DB_PARAMS)" down -all

auth/migrate/create:
	$(MIGRATE_CREATE) $(AUTH_MIGRATIONS) $(name)

### Matcher service targets
matcher/check-coverage:
	-cd matcher && go test ./... -coverprofile=cover.out -covermode=atomic -coverpkg=./...
	-cd matcher && go-test-coverage --config=.local.testcoverage.yml
	@rm matcher/cover.out

matcher/local-build:
	cd matcher/cmd/matcher && go build -ldflags="-s -w"

matcher/lint:
	cd matcher && golangci-lint run

matcher/run:
	docker compose up -d matcher

matcher/start: matcher/generate
	docker compose up --build --force-recreate -d matcher

matcher/generate:
	cd matcher && go generate ./...

MATCHER_MIGRATIONS = file://matcher/migrations
name = create_filter_interests_table

matcher/migrate/up:
	migrate -verbose -source $(MATCHER_MIGRATIONS) -database "$(DB)matcher$(DB_PARAMS)" up

matcher/migrate/down:
	migrate -source $(MATCHER_MIGRATIONS) -database "$(DB)matcher$(DB_PARAMS)" down -all

matcher/migrate/create:
	$(MIGRATE_CREATE) $(MATCHER_MIGRATIONS) $(name)