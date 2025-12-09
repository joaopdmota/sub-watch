.PHONY: up
up:
	 docker compose up -d

.PHONY: down	
down:
	 docker compose down

.PHONY: server-logs
server-logs:
	 docker compose logs -f bff

.PHONY: install
install:
	 docker compose run --rm bff npm install