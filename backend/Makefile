.PHONY: build
build:
	@echo "Building..."
	docker compose build --no-cache
	@echo "Building finished."

.PHONY: up
up:
	@echo "Initializing..."
	docker compose up -d
	@echo "Initialization finished."

.PHONY: down
down:
	@echo "Killing..."
	docker compose down
	@echo "Kill finished."

.PHONY: test
test:
	@echo "Running tests..."
	go test $(shell go list ./... | grep -v '/mocks') -coverprofile=coverage.out
	go tool cover -func=coverage.out
	@echo "Tests completed."

.PHONY: tidy
tidy:
	@echo "Tidying go modules inside container..."
	docker compose run --rm server go mod tidy
	@echo "Modules tidied."


.PHONY: test-cover
test-cover:
	@echo "Running tests coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Tests completed."

.PHONY: run
run:
	@echo "Running app..."
	docker compose up -d
	@echo "Build completed."

.PHONY: build-docker
build-docker:
	@echo "Running build..."
	docker build --build-arg WEATHER_API_KEY=$(WEATHER_API_KEY) --build-arg API_PORT=$(API_PORT) -f $(DOCKERFILE) -t $(IMAGE_NAME) . --no-cache
	@echo "Build completed."

.PHONY: run-docker-build
run-docker-build:
	@echo "Running app..."
	docker run -e WEATHER_API_KEY=$(WEATHER_API_KEY) -e API_PORT=$(API_PORT) -p $(API_PORT):$(API_PORT) cep_weather

.PHONE: push
push: build-docker
	@echo "Pushing image..."
	docker push $(DOCKER_USERNAME)/$(IMAGE_NAME)
	@echo "Push completed."

.PHONY: env
env:
	@echo "Copying .env.example to .env..."
	cp .env.example .env
	@echo ".env file created."

.PHONY: server-logs
server-logs:
	@echo "Tailing server logs..."
	docker compose logs -f server
	@echo "Logs streaming stopped."