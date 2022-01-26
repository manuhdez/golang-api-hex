up:
	@docker compose up -d
build:
	@docker compose build --no-cache
down:
	@docker compose down
start:
	@docker compose start

test:
	go test ./...
