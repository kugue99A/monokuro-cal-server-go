.PHONY: setup dev build sqlc migrate rollback db-up db-down

setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

dev:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

sqlc:
	sqlc generate

migrate:
	goose -dir db/migrations postgres "$(DATABASE_URL)" up

rollback:
	goose -dir db/migrations postgres "$(DATABASE_URL)" down

db-up:
	docker compose up -d

db-down:
	docker compose down
