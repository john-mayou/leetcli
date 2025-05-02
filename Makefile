-include .env
export

OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

# ----- Deps -----

GOLANGCI_LINT_VERSION = 2.1.5
GOLANG_MIGRATE_VERSION = 4.18.3

check-docker:
	@command -v docker >/dev/null 2>&1 || \
	{ echo >&2 "'docker' is not installed. Please install Docker: https://docs.docker.com/get-docker/"; exit 1; }

check-compose:
	@docker compose version >/dev/null 2>&1 || \
	{ echo >&2 "'docker compose' is not available. Please install Docker Compose V2: https://docs.docker.com/compose/install/"; exit 1; }

check-env-database-url:
	@test -n "$$DATABASE_URL" || (echo "DATABASE_URL is not set"; exit 1)

check-env-migrate-force-version:
	@test -n "$$MIGRATE_FORCE_VERSION" || (echo "MIGRATE_FORCE_VERSION is not set"; exit 1)

install-golangci:
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION); \
	}

install-golang-migrate:
	@if ! command -v migrate >/dev/null 2>&1 && [ ! -x .tmp/golang-migrate/migrate ]; then \
		echo "Installing golang-migrate $(GOLANG_MIGRATE_VERSION)..."; \
		mkdir -p ./.tmp/golang-migrate; \
		curl -sSL https://github.com/golang-migrate/migrate/releases/download/v$(GOLANG_MIGRATE_VERSION)/migrate.$(OS)-$(ARCH).tar.gz | tar xz -C ./.tmp/golang-migrate; \
	fi

# ----- Docker Compose -----

COMPOSE_DEV = docker compose -f docker-compose.dev.yml
COMPOSE_TEST = docker compose -f docker-compose.test.yml

dev-api-up:
	$(COMPOSE_DEV) --profile api up -d --build

dev-db-logs:
	$(COMPOSE_DEV) logs --follow db

dev-api-logs:
	$(COMPOSE_DEV) logs --follow api

dev-api-down:
	$(COMPOSE_DEV) --profile api down

# ----- Testing -----

test-api: check-docker check-compose
	set -exuo pipefail; \
	$(COMPOSE_TEST) --profile api up -d --build; \
	trap "$(COMPOSE_TEST) --profile api down" EXIT; \
	$(COMPOSE_TEST) exec api sh -c "cd api && go test ./... -v"

# ----- Linting -----

lint-api: install-golangci
	@cd api && golangci-lint run

# ----- Migration / DB -----

MIGRATE = $(shell command -v migrate 2>/dev/null || echo .tmp/golang-migrate/migrate)

db-migrate-up: check-env-database-url install-golang-migrate
	$(MIGRATE) -source file://migrations -database "$$DATABASE_URL" up

db-migrate-down: check-env-database-url install-golang-migrate
	$(MIGRATE) -source file://migrations -database "$$DATABASE_URL" down 1

db-migrate-force: check-env-migrate-force-version check-env-database-url install-golang-migrate
	$(MIGRATE) -source file://migrations -database "$$DATABASE_URL" force $$MIGRATE_FORCE_VERSION