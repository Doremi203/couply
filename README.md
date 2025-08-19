# Couply Monorepo

Monorepo for a full-stack application with Go microservices, a React + TypeScript + Vite frontend, and auxiliary Go cloud functions.

- Backend: Go (gRPC + gRPC-Gateway), PostgreSQL, S3-compatible storage (MinIO), Docker Compose
- Frontend: React 18, TypeScript, Vite, Storybook, Vitest, E2E via Testplane/Playwright
- Cloud functions: Go cloud functions for photo upload and verification

## Repository structure

- backend/ — Go microservices, shared protos, compose, migrations, tests
  - auth/ — authentication service
  - matcher/ — matching service
  - blocker/ — blocking service
  - notificator/ — notifications service
  - payments/ — payments service
  - common/ — shared libraries, proto definitions
  - docker-compose.yml — local stack (Postgres, MinIO, all services)
  - Makefile — dev tooling, codegen, migrations, lint, test
- frontend/ — React + TS + Vite app, Storybook, tests, E2E
- cloud-functions/ — Go cloud functions
  - photo-loader — S3 photo upload function
  - verifier/ — ML photo analysis function

## Tech stack

- Go 1.24+ (backend), Go 1.21+ (cloud-functions)
- gRPC + gRPC-Gateway, Protocol Buffers powered by Buf
- PostgreSQL , migrations via goose
- S3 for object storage
- Valkey for cahching and rate limiting
- React 18, TypeScript, Vite, TailwindCSS, MUI
- Testing: Go test + coverage, Vitest/Storybook tests, Testplane (E2E) with Playwright
- Linting: golangci-lint (Go), ESLint + Prettier (Web)

## Testing and linting

Backend (from backend/):

- Lint all: make lint
- Service coverage: make auth/check-coverage (and matcher/blocker/notificator/payments)
- Integration compose (if used): docker compose -f docker-compose.tests.yml up

Frontend (from frontend/):

- Type/lint/format check: npm run check
- Unit/Component tests: npm run test-storybook (with coverage via npm run storybook-coverage)
- Storybook: npm run storybook
- E2E (GUI): npm run e2e (Testplane + Playwright)

## Code generation (Protobuf/Buf)

- Tools install: make deps (includes buf, protoc plugins)
- Pull external protos:
  - make common/pull/proto (and specific targets for googleapis, openapiv2, validate)
- Generate all services: make generate
- Generate per service: make auth/generate (and matcher/blocker/notificator/payments)

Generated artifacts include:

- Protobuf Go messages, gRPC services, gRPC-Gateway, OpenAPI, validation code

## Environment and configuration

- Backend compose environment is defined under backend/docker-compose.yml
- Secrets (local): backend/secrets.env
  - You can generate a YC token into secrets.env: make secrets/yc-token
- Other secrets are pulled from Yandex Cloud Secrets Manager
- Services look for configs at CONFIGS_PATH=/app/configs (default in compose)
