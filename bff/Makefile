-include .env

SERVICE_NAME ?= sub-watch-bff
SERVICE_PORT ?= 3001
IMAGE_NAME ?= $(SERVICE_NAME)
DOCKERFILE ?= Dockerfile.dev

COMPOSE := docker compose

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        Build docker images (no cache)"
	@echo "  make up           Start containers"
	@echo "  make down         Stop containers"
	@echo "  make restart      Restart containers"
	@echo "  make logs         Tail application logs"
	@echo "  make shell        Access container shell"
	@echo "  make install      Install npm dependencies inside container"
	@echo "  make test         Run tests (inside container)"
	@echo "  make lint         Run linter (if configured)"
	@echo "  make env          Create .env file from example"
	@echo "  make clean        Remove containers, volumes and cache"

# ---------------------------------------------------
# Docker / Compose
# ---------------------------------------------------

.PHONY: build
build:
	@echo "Building images..."
	$(COMPOSE) build --no-cache
	@echo "Build finished."

.PHONY: up
up:
	@echo "Starting containers..."
	$(COMPOSE) up -d
	@echo "Containers started."

.PHONY: down
down:
	@echo "Stopping containers..."
	$(COMPOSE) down
	@echo "Containers stopped."

.PHONY: restart
restart: down up

.PHONY: clean
clean:
	@echo "Cleaning containers, volumes and cache..."
	$(COMPOSE) down -v --remove-orphans
	docker system prune -f
	@echo "Clean finished."

# ---------------------------------------------------
# Dev & Debug
# ---------------------------------------------------

.PHONY: logs
logs:
	@echo "Tailing logs..."
	$(COMPOSE) logs -f bff

.PHONY: shell
shell:
	@echo "Opening shell inside container..."
	$(COMPOSE) exec bff sh

.PHONY: install
install:
	@echo "Installing dependencies inside container..."
	$(COMPOSE) run --rm bff npm install
	@echo "Dependencies installed."

# ---------------------------------------------------
# App lifecycle
# ---------------------------------------------------

.PHONY: run
run:
	@echo "Running app..."
	$(COMPOSE) up
	@echo "Run finished."

# ---------------------------------------------------
# Quality
# ---------------------------------------------------

.PHONY: test
test:
	@echo "Running tests..."
	$(COMPOSE) run --rm bff npm test
	@echo "Tests completed."

.PHONY: lint
lint:
	@echo "Running linter..."
	$(COMPOSE) run --rm bff npm run lint
	@echo "Lint completed."

# ---------------------------------------------------
# Environment
# ---------------------------------------------------

.PHONY: env
env:
	@echo "Creating .env from example..."
	cp .env.example .env
	@echo ".env created."
