.PHONY: help build run test clean install migrate-up migrate-down docker-up docker-down

APP_NAME=boilerplate-go
BINARY_NAME=main
DOCKER_IMAGE=$(APP_NAME)
DOCKER_CONTAINER=$(APP_NAME)_app

# Variables
MAIN_PATH=cmd/api

# Load .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@CGO_ENABLED=0 go build -o $(BINARY_NAME) ./$(MAIN_PATH)
	@echo "Build complete: $(BINARY_NAME)"

run: ## Run the application
	@echo "Running $(APP_NAME)..."
	@go run ./$(MAIN_PATH)/main.go

dev: ## Run application in development mode with hot reload
	@echo "Running $(APP_NAME) in development mode..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "air not installed. Install it with: go install github.com/cosmtrek/air@latest"; \
		go run ./$(MAIN_PATH)/main.go; \
	fi

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install it from: https://golangci-lint.run/usage/install/"; \
	fi

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

clean: ## Clean build artifacts
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@go clean

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image: $(DOCKER_IMAGE)"
	@docker build -t $(DOCKER_IMAGE):latest .

docker-run: docker-build ## Run Docker container
	@echo "Running Docker container: $(DOCKER_CONTAINER)"
	@docker run -d \
		--name $(DOCKER_CONTAINER) \
		--network boilerplate_network \
		-p 8080:8080 \
		--env-file .env \
		$(DOCKER_IMAGE):latest

docker-stop: ## Stop Docker container
	@echo "Stopping Docker container: $(DOCKER_CONTAINER)"
	@docker stop $(DOCKER_CONTAINER)
	@docker rm $(DOCKER_CONTAINER)

# Docker Compose commands
docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose up -d

docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services with docker-compose..."
	@docker-compose down -v

docker-compose-logs: ## View docker-compose logs
	@docker-compose logs -f

docker-compose-restart: ## Restart services with docker-compose
	@echo "Restarting services with docker-compose..."
	@docker-compose restart

# Migration commands
migrate-create-%: ## Create a new migration file. Usage: make migrate-create-users_table
	@echo "Creating migration: $*"
	@mkdir -p migrations
	@touch migrations/$$(date +%s)_$*.up.sql
	@touch migrations/$$(date +%s)_$*.down.sql

migrate-up: ## Run all migrations
	@echo "Running migrations up..."
	@if command -v migrate > /dev/null; then \
		migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up; \
	else \
		echo "migrate not installed. Install it with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-down: ## Rollback migrations
	@echo "Running migrations down..."
	@if command -v migrate > /dev/null; then \
		migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down; \
	else \
		echo "migrate not installed. Install it with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-version: ## Check migration version
	@if command -v migrate > /dev/null; then \
		migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" version; \
	else \
		echo "migrate not installed."; \
	fi

# Development setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	@make install
	@echo "Creating .env file from .env.example..."
	@cp .env.example .env
	@echo "Setup complete! Run 'make dev' to start the application."

all: install build ## Install dependencies and build

.DEFAULT_GOAL := help
