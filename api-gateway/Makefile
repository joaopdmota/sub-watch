PROJECT_NAME := sub-watch-api-gateway
DOCKER_COMPOSE_FILE := docker-compose.yaml
KRAKEND_CONFIG := krakend.json
DOCKER_EXEC := docker compose exec $(PROJECT_NAME)

.PHONY: up
up:
	@echo "🚀 Starting API Gateway $(PROJECT_NAME)..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: down
down:
	@echo "🛑 Stopping and removing containers of $(PROJECT_NAME)..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: rebuild
rebuild: down
	@echo "🏗️ Rebuilding and starting the service..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d --build

.PHONY: restart
restart:
	@echo "🔄 Restarting the container of $(PROJECT_NAME)..."
	docker compose -f $(DOCKER_COMPOSE_FILE) restart krakend-gateway

.PHONY: logs
logs:
	@echo "📄 Showing logs of $(PROJECT_NAME)... (Pressione Ctrl+C para sair)"
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f krakend-gateway

.PHONY: check-config
check-config:
	@echo "✅ Validating the syntax of $(KRAKEND_CONFIG)..."
	$(DOCKER_EXEC) krakend check -c /etc/krakend/$(KRAKEND_CONFIG)