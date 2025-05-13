# Backend Services

This README provides instructions for setting up and running the backend services for local development.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Development Setup](#development-setup)
- [Running Services](#running-services)
- [Database Migrations](#database-migrations)
- [Testing](#testing)
- [Linting](#linting)
- [Code Generation](#code-generation)

## Prerequisites

The following tools are required to run the backend services:

- Go 1.24+
- Docker and Docker Compose
- PostgreSQL
- [Goose](https://github.com/pressly/goose) (for migrations)
- [Yandex Cloud CLI](https://cloud.yandex.com/en/docs/cli/quickstart) (for token generation)
- golangci-lint (for code linting)

## Project Structure

The project consists of several microservices:
- **auth**: Authentication service
- **matcher**: Matching service
- **blocker**: Blocking service
- **notificator**: Notification service
- **common**: Shared code and protocols

## Development Setup

1. **Install dependencies**:
```shell script
make deps
```


This will install all required development tools, including:
- Migration tools (goose)
- Test coverage tools
- Protocol buffer tools (protoc, buf)
- Mock generators
- Linters

2. **Generate Yandex Cloud token** (if required):
```shell script
make secrets/yc-token
```


3. **Generate code** for all services:
```shell script
make generate
```


## Running Services

## Docker Compose Configuration

The project uses Docker Compose for local development and testing. Understanding the different Docker Compose configurations will help you efficiently develop and debug the services.

### Main Docker Compose Setup

The primary Docker Compose configuration is defined in `docker-compose.yml` at the project root. This file:

- Defines all microservices (auth, matcher, blocker, notificator)
- Sets up the PostgreSQL database service
- Configures networks and volumes
- Sets environment variables for services
- Maps service ports to the host machine

To start all services with the main configuration:
```shell script
docker compose up
```

Or to run in detached mode:
```shell script
docker compose up -d
```

### Debug Configurations

Each service has its own debug configuration in its directory (`auth/docker-compose.debug.yml`, `matcher/docker-compose.debug.yml`, etc.). These configurations:

- Override settings from the main docker-compose.yml
- Set up services with debug ports exposed
- Configure Delve debugger for Go services
- Allow you to connect your IDE to the running service for debugging

To run a service in debug mode:
```shell script
# For the auth service
cd auth
docker compose -f ../docker-compose.yml -f docker-compose.debug.yml up auth

# For other services
cd matcher
docker compose -f ../docker-compose.yml -f docker-compose.debug.yml up matcher
```

### Remote Debugging

To debug a service:

1. Start the service in debug mode using the appropriate docker-compose.debug.yml file
2. Configure your IDE to connect to the debug port (typically 40000+ range)
3. Set breakpoints in your code
4. Execute the code path that triggers the breakpoint

Most JetBrains IDEs (like GoLand) support Go remote debugging out of the box. Configure a "Go Remote" run configuration with the appropriate host and port.

### Testing Configuration

The project includes testing configurations:
- `docker-compose.tests.yml`: Used for running integration tests
- `docker-compose.testing-template.yml`: Template for setting up test environments

To run the test suite:
```shell script
docker compose -f docker-compose.tests.yml up
```

### Environment Variables

Environment variables for services are defined in the Docker Compose files. Common variables include:
- Database connection parameters
- Service API endpoints
- Authentication settings
- Feature flags
- Log levels

If you need to override environment variables, you can:
1. Modify the docker-compose.yml file directly
2. Create a docker-compose.override.yml file with your custom settings
3. Pass environment variables on the command line:
```shell script
LOG_LEVEL=debug docker compose up
```

### Service Dependencies

The Docker Compose configurations handle service dependencies, ensuring that:
- The database starts before the services
- Services with dependencies start in the correct order
- Health checks verify that services are ready

If you need to run a specific service with its dependencies:
```shell script
docker compose up postgres auth
```

This will start only the PostgreSQL database and the auth service.


## Database Migrations

Each service has its own database schema and migration commands.

### Auth Service

```shell script
# Apply migrations
make auth/migrate/up

# Revert all migrations
make auth/migrate/down

# Revert single migration
make auth/migrate/down-one

# Create new migration
make auth/migrate/create name=migration_name
```


### Matcher Service

```shell script
# Apply migrations
make matcher/migrate/up

# Revert all migrations
make matcher/migrate/down

# Revert single migration
make matcher/migrate/down-one

# Create new migration
make matcher/migrate/create name=migration_name
```


### Blocker Service

```shell script
# Apply migrations
make blocker/migrate/up

# Revert all migrations
make blocker/migrate/down

# Revert single migration
make blocker/migrate/down-one

# Create new migration
make blocker/migrate/create name=migration_name
```


### Notificator Service

```shell script
# Apply migrations
make notificator/migrate/up

# Revert all migrations
make notificator/migrate/down

# Revert single migration
make notificator/migrate/down-one

# Create new migration
make notificator/migrate/create name=migration_name
```

## Testing

The project uses Go's testing tools with coverage reporting:

```shell script
# Check test coverage for auth service
make auth/check-coverage

# Check test coverage for matcher service
make matcher/check-coverage

# Check test coverage for blocker service
make blocker/check-coverage

# Check test coverage for notificator service
make notificator/check-coverage
```


## Linting

Run linting for all services:
```shell script
make lint
```


Or for specific services:
```shell script
make auth/lint
make matcher/lint
make blocker/lint
```


## Local Builds

To verify that services build successfully without running them:

```shell script
make auth/local-build
make matcher/local-build
make blocker/local-build
make notificator/local-build
```

## Development Environment

The default database configuration is:
- Host: localhost
- Port: 5432
- Username: user
- Password: pass
- Database names: auth, matcher, blocker, notificator

These settings can be modified in the Makefile if needed.

## Code Generation

Code generation is used for protocol buffers, mocks, and other generated code:

```shell script
# Generate code for all services
make generate

# Generate code for specific service
make auth/generate
make matcher/generate
make blocker/generate
make notificator/generate
```

## Protocol Buffer Code Generation with Buf

The project uses [Buf](https://buf.build/) for managing Protocol Buffer definitions and generating code. This section explains how the Buf configuration is set up and how to work with it.

## Buf Configuration Files

The Buf setup consists of two main configuration files in the project root:

1. **buf.yaml** - The main Buf configuration file that defines:
    - Protocol buffer version
    - Lint and breaking change detection rules
    - Dependencies on external Protocol Buffer repositories

2. **buf.gen.yaml** - This file defines code generation options:
    - Which languages to generate (Go in this case)
    - Plugin options for gRPC, validation, and API gateway
    - Output directories for generated code

## Protocol Buffer Structure

Protocol Buffer definitions are stored in the `/common/proto` directory, which contains:
- Service definitions
- Message definitions
- API gateway configurations
- Validation rules

The project also pulls in external Protocol Buffer definitions from:
- Google APIs (HTTP annotations, field behaviors)
- gRPC Gateway (OpenAPI annotations)
- Protoc-gen-validate (validation rules)

## Generating Protocol Buffer Code

The code generation process uses these make commands:

```shell script
# Install Buf and related tools
make deps/buf
make deps/protoc

# Pull external Protocol Buffer definitions
make common/pull/proto
make common/pull/proto/googleapis
make common/pull/proto/protoc-gen-openapiv2
make common/pull/proto/validate

# Generate code for all services (includes Protocol Buffer generation)
make generate

# Generate code for a specific service
make auth/generate
make matcher/generate
make blocker/generate
make notificator/generate
```


## Generated Artifacts

Buf generates several types of Go code:
- **Plain Protocol Buffer Go code**: Message definitions, enums, etc.
- **gRPC service definitions**: Service interfaces and client implementations
- **gRPC Gateway code**: HTTP/JSON gateway for gRPC services
- **OpenAPI definitions**: Swagger/OpenAPI documentation
- **Validation code**: Request validation based on annotations

The generated code is automatically included in the Go builds and doesn't need to be committed to version control.

## Modifying Protocol Definitions

When modifying Protocol Buffer definitions:

1. Edit the `.proto` files in the `/common/proto` directory
2. Run `make generate` to regenerate code
3. For breaking changes, follow the backward compatibility guidelines enforced by Buf

Remember that Protocol Buffers form the contract between services, so changes should be made carefully to maintain compatibility.