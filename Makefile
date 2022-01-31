up:
	@docker compose up -d
.PHONY: build
build:
	@docker compose build --no-cache
down:
	@docker compose down
start:
	@docker compose start

test:
	go test ./...

deps:
	go mod download
	go mod verify
	@make tidy

tidy:
	go mod tidy
