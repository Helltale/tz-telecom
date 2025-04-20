# ========== CONFIGURATION ==========
APP_NAME=tz-telecom
GO_CMD=go
MAIN_PACKAGE=./cmd/main.go
ENV_FILE=.env

export APP_PORT ?= 8080
export DB_HOST ?= localhost
export DB_PORT ?= 5432
export DB_USER ?= postgres
export DB_PASS ?= postgres
export DB_NAME ?= tz_telecom
export DB_SSLMODE ?= disable
export APP_READ_TIMEOUT ?= 10
export APP_WRITE_TIMEOUT ?= 10
export APP_IDLE_TIMEOUT ?= 120
export APP_WORKER_QUEUE_LEN ?= 100
export DB_RETRY_INITIAL_DELAY ?= 1
export DB_RETRY_MAX_DELAY ?= 10
export DB_RETRY_MULTIPLIER ?= 2
export DB_RETRY_MAX_ATTEMPTS ?= 5

# ========== COMMANDS ==========

serve:
	@echo "starting $(APP_NAME)..."
	$(GO_CMD) run $(MAIN_PACKAGE) serve

migrate:
	@echo "running migrations..."
	$(GO_CMD) run $(MAIN_PACKAGE) migrate

deps:
	@echo "installing dependencies..."
	$(GO_CMD) mod tidy

test:
	@echo "running tests..."
	$(GO_CMD) test -v ./...

fmt:
	@echo "formatting code..."
	$(GO_CMD) fmt ./...

env:
	@echo "APP_PORT=$(APP_PORT)"
	@echo "DB_HOST=$(DB_HOST)"
	@echo "DB_PORT=$(DB_PORT)"
	@echo "DB_USER=$(DB_USER)"
	@echo "DB_PASS=$(DB_PASS)"
	@echo "DB_NAME=$(DB_NAME)"
	@echo "DB_SSLMODE=$(DB_SSLMODE)"
	@echo "APP_READ_TIMEOUT=$(APP_READ_TIMEOUT)"
	@echo "APP_WRITE_TIMEOUT=$(APP_WRITE_TIMEOUT)"
	@echo "APP_IDLE_TIMEOUT=$(APP_IDLE_TIMEOUT)"
	@echo "APP_WORKER_QUEUE_LEN=$(APP_WORKER_QUEUE_LEN)"
	@echo "DB_RETRY_INITIAL_DELAY=$(DB_RETRY_INITIAL_DELAY)"
	@echo "DB_RETRY_MAX_DELAY=$(DB_RETRY_MAX_DELAY)"
	@echo "DB_RETRY_MULTIPLIER=$(DB_RETRY_MULTIPLIER)"
	@echo "DB_RETRY_MAX_ATTEMPTS=$(DB_RETRY_MAX_ATTEMPTS)"

run: deps migrate serve
